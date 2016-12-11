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

func GetCacheFileName(where, what string) (string, error) {
	cache_dir := HomeDotDir + "/dumps/" + where + "/"
	if exists, _ := DirExists(cache_dir); !exists {
		if err := CreatePath(cache_dir); err != nil {
			return "", err
		}
	}
	filename := cache_dir
	switch where {
	case "local":
		_, dir_name := path.Split(what)
		filename += dir_name
	case "remote":
		filename += what
	default:
		panic("Unknown option: " + where)
	}
	return filename + ".json", nil
}

func CacheFileList(fileList FileList, where, what string) {
	b, err := json.Marshal(fileList)
	if err != nil {
		Error.Printf("Error %s.", err)
	}
	filename, err := GetCacheFileName(where, what)
	if err != nil {
		Error.Println(err.Error())
		return
	}
	ioutil.WriteFile(filename, b, 0644)
	Debug.Printf("Dumping the filelist into the file: %s.", filename)
}

func LoadFileTreeFromCache(where, what string) (*FileList, error) {
	var fileTree FileList
	filename, err := GetCacheFileName(where, what)
	if err != nil {
		return nil, err
	}
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
