package cache

import (
	"os"
	"path"
	"errors"
	"encoding/json"
	"io/ioutil"

	. "github.com/taryk/gdtool/core"
)

func CheckCacheExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func GetCacheFileName(where, what string) string {
	filename := DotFolder + "/dumps/" + where + "/"
	switch where {
	case "local":
		_, dir_name := path.Split(what)
		filename += dir_name
	case "remote":
		filename += what
	default:
		panic("Unknown option: " + where)
	}
	return filename + ".json"
}

func CacheFileList(fileList FileList, where, what string) {
	b, err := json.Marshal(fileList)
	if err != nil {
		Error.Printf("Error %s.", err)
	}
	filename := GetCacheFileName(where, what)
	ioutil.WriteFile(filename, b, 0644)
	Debug.Printf("Dumping the filelist into the file: %s.", filename)
}

func LoadFileTreeFromCache(where, what string) (*FileList, error) {
	var fileTree FileList
	filename := GetCacheFileName(where, what)
	if !CheckCacheExists(filename) {
		return nil, errors.New("File " + filename + " doesn't exist")
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		Error.Printf("Error %s", err)
		return nil, err
	}
	if err := json.Unmarshal(b, &fileTree); err != nil {
		Error.Printf("Error %s", err)
		return nil, err
	}
	return &fileTree, nil
}
