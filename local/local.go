package local

import (
	"io"
	"os"
	"crypto/md5"
	"encoding/hex"

	. "github.com/taryk/gdtool/core"
)

func getMD5sum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		Error.Printf("Error when opening a file: \"%s\": %s.", filename,
			err.Error())
		return "", err
	}
	defer f.Close()
	st, err := f.Stat()
	if err != nil {
		Error.Printf("Error when getting file details \"%s\": %s.", filename,
			err.Error())
		return "", err
	}
	if st.IsDir() {
		Error.Printf("\"%s\" is a directory.", filename)
		return "", err
	}
	md5 := md5.New()
	io.Copy(md5, f)
	return hex.EncodeToString(md5.Sum(nil)), nil
}

func getFileDetails(absPath string, fi os.FileInfo) FileDetails {
	md5sum, _ := getMD5sum(absPath + "/" + fi.Name())
	var children FileList
	return FileDetails{
		Id:       "",
		Name:     fi.Name(),
		Size:     uint64(fi.Size()),
		Md5sum:   md5sum,
		IsDir:    fi.IsDir(),
		Children: children,
	}
}

func GetFileList(absPath string, recursive bool) (FileList, error) {
	dir, err := os.Open(absPath)
	if err != nil {
		Error.Printf("Error when opening the directory \"%s\": %s.", absPath,
			err)
		return nil, err
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		Error.Printf("Error when reading the directory content \"%s\": %s.",
			absPath, err)
		return nil, err
	}
	var result FileList
	for _, fi := range fileInfos {
		fileDetails := getFileDetails(absPath, fi)
		if recursive && fileDetails.IsDir {
			fileDetails.Children, _ = GetFileList(
				absPath + "/" + fileDetails.Name, recursive)
		}
		result = append(result, &fileDetails)
	}
	return result, nil
}
