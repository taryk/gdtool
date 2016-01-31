package remote

import (
	"time"

	"google.golang.org/api/drive/v2"

	. "github.com/taryk/gdtool/core"
)

type FileDetailsResult struct {
	FileDetails *FileDetails
	Position int64
}

// Number of requests per second allowed.
const req_per_sec int = 10

var c_request_no chan int64 = make(chan int64)
var c_last_request_time chan time.Time = make(chan time.Time)

func Init() {
	go requestTimeHandler()
}

func checkRequestLimit() {
	request_no := getRequestNo()
	request_no++
	if request_no % int64(req_per_sec) == 0 {
		dur := time.Since(getLastRequestTime())
		if dur.Seconds() < 1 {
			time_to_sleep := 1 * time.Second -
				time.Duration(dur.Nanoseconds())
			time.Sleep(time_to_sleep)
		}
	}
	// We ignore what we send back.
	c_request_no <- request_no
}

func getLastRequestTime() time.Time {
	return <- c_last_request_time
}

func getRequestNo() int64 {
	return <- c_request_no
}

func updateLastRequestTime() {
	if getRequestNo() % int64(req_per_sec) == 0 {
		c_last_request_time <- time.Now()
	}
}

func requestTimeHandler() {
	// Time since the latest request was made.
	var last_request_time time.Time = time.Now()

	// Number of requests made. We've already done 2 requests so far.
	var request_no int64 = 2

	for {
		select {
		case last_request_time = <- c_last_request_time:
			// Do nothing.
			Debug.Printf("Update last_request_time %+v.", last_request_time)
		case c_last_request_time <- last_request_time:
			Debug.Printf("Send last_request_time %+v.", last_request_time)
		case <- c_request_no:
			// Ignore the received value, just increment request_no.
			request_no++
			Debug.Printf("Update request_no %d.", request_no)
		case c_request_no <- request_no:
			Debug.Printf("Send request_no %d.", request_no)
		case <- time.After(time.Second):
			// Do nothing.
			Debug.Print("Timeout.")
		}
	}
}

// Fetches all children of a given folder.
func allChildren(d *drive.Service, folderId string) ([]*drive.ChildReference,
	error) {

	var cs []*drive.ChildReference
	pageToken := ""
	for {
		q := d.Children.List(folderId)
		// If we have a pageToken set, apply it to the query.
		if pageToken != "" {
			q = q.PageToken(pageToken)
		}
		var r *drive.ChildList
		var err error
	request:
		for {
			// Check if we don't exceed the request limit.
			checkRequestLimit()
			r, err = q.Do()
			if err != nil {
				Error.Printf("An error occurred: %v.", err)
				time.Sleep(100 * time.Millisecond)
				continue request
			}
			updateLastRequestTime()
			break request
		}
		cs = append(cs, r.Items...)
		pageToken = r.NextPageToken
		if pageToken == "" {
			break
		}
	}
	return cs, nil
}

func isDir(f *drive.File) bool {
	return f.MimeType == "application/vnd.google-apps.folder"
}

func getFileDetails(c chan FileDetailsResult, d *drive.Service, position int64,
	fileId string) {

	result := FileDetailsResult{
		Position: position,
	}
	checkRequestLimit()
	f, err := d.Files.Get(fileId).Do()
	updateLastRequestTime()
	if err != nil {
		Error.Printf("An error occurred: %v.", err)
		c <- result
		return
	}	
	var children FileList

	result.FileDetails = &FileDetails{
		Id:       fileId,
		Name:     f.Title,
		Size:     uint64(f.FileSize),
		Md5sum:   f.Md5Checksum,
		IsDir:    isDir(f),
		Children: children,
	}
	c <- result
}

func receiveData(n int, c chan FileDetailsResult, fileList *FileList,
	failedPositions *[]int64) {

	for i := 1; i <= n; i++ {
		Debug.Printf("Expect to receive = %d/%d.", i, n)
		fileDetailsResult := <-c
		(*fileList)[fileDetailsResult.Position] = fileDetailsResult.FileDetails
		if fileDetailsResult.FileDetails == nil {
			*failedPositions = append(
				*failedPositions, fileDetailsResult.Position)
		}
		Debug.Printf("Received = %d/%d.", i, n)
	}
}

func GetFileList(d *drive.Service, folderId string, recursive bool) (
	FileList, error) {

	list, err := allChildren(d, folderId)
	if err != nil {
		Error.Fatalf("Error occured: %v.", err)
		return nil, err
	}
	result := make(FileList, len(list))
	c := make(chan FileDetailsResult)

	var max_workers int
	request_no := getRequestNo()
	if request_no < int64(req_per_sec) {
		max_workers = req_per_sec - int(request_no)
	} else {
		max_workers = req_per_sec
	}
	Debug.Printf("Max_workers (initial) = %d / requests = %d.",
		max_workers, request_no)
	n := 0
	var failedPositions []int64
	for position, v := range list {
		go getFileDetails(c, d, int64(position), v.Id)
		n++
		Debug.Printf("Run goroutine n = %d | position = %d.", n, position)
		// Collect data every so often.
		if n >= max_workers {
			receiveData(n, c, &result, &failedPositions)
			if max_workers != req_per_sec {
				max_workers = req_per_sec
				Debug.Printf("Max_workers new = %d.", max_workers)
			}
			n = 0
		}
	}
	if n > 0 {
		receiveData(n, c, &result, &failedPositions)
		n = 0
	}
	// Process failed requests.
	for len(failedPositions) > 0 {
		Debug.Printf("There are %d failed requests.", len(failedPositions))
		var newFailedPositions []int64
		for _, position := range failedPositions {
			go getFileDetails(c, d, position, list[position].Id)
			n++
			Debug.Printf("Repair - run goroutine n = %d / position = %d.", n,
				position)
			if n >= max_workers {
				receiveData(n, c, &result, &newFailedPositions)
				if max_workers != req_per_sec {
					max_workers = req_per_sec
					Debug.Printf("Repair - max_workers new = %d.", max_workers)
				}
				n = 0
			}
		}
		if n > 0 {
			receiveData(n, c, &result, &newFailedPositions)
			n = 0
		}
		failedPositions = newFailedPositions
	}
	Debug.Printf("Finished. Req no = %d.", getRequestNo())
	close(c)
	for _, fileDetails := range result {
		if recursive && fileDetails.IsDir {
			fileDetails.Children, _ = GetFileList(d, fileDetails.Id, recursive)
		}
	}
	return result, nil
}
