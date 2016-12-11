package core

import (
	"fmt"
	"log"
	"sort"
	"os"
	"os/user"
	"io"
	"strings"
	"path/filepath"

	"github.com/fatih/color"
)

var usr, _ = user.Current()
var HomeDotDir = filepath.Join(usr.HomeDir, ".gdtool")
var TokenCacheDir = filepath.Join(HomeDotDir, "credentials")

var IsTesting bool = false

var (
	Debug *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

type FileDetails struct {
	Id, Name, Path string
	Size           uint64
	Md5sum         string
	IsDir          bool
	Children       []*FileDetails
}

type FileList []*FileDetails

type FileDetailsMap struct {
	Id, Name, Path string
	Size           uint64
	Md5sum         string
	IsDir          bool
	Children       map[string][]*FileDetailsMap
}

type FileTreeMap map[string][]*FileDetailsMap

// sort by name
type ByName FileList

func (fd ByName) Len() int {
	return len(fd)
}

func (fd ByName) Less(i, j int) bool {
	return fd[i].Name < fd[j].Name
}

func (fd ByName) Swap(i, j int) {
	fd[i], fd[j] = fd[j], fd[i]
}

//

func InitLoggers(log_files ...string) {
	loggers := map[string]**log.Logger{
		"debug": &Debug,
		"warn":  &Warn,
		"error": &Error,
	}
	handler:
	for _, name := range log_files {
		if _, ok := loggers[name]; !ok {
			continue handler
		}
		var log_handler io.Writer
		if IsTesting {
			log_handler = os.Stdout
		} else {
			log_dir := HomeDotDir + "/logs"
			if exists, _ := DirExists(log_dir); !exists {
				if err := CreatePath(log_dir); err != nil {
					fmt.Errorf("Can't create a log dir \"%s\"\n", log_dir)
					continue handler
				}
			}
			log_handler, _ = os.OpenFile(log_dir + "/" + name + ".log",
				os.O_CREATE | os.O_RDWR | os.O_APPEND, 0660)
		}
		*loggers[name] = log.New(log_handler,
			strings.ToUpper(name) + ": ",
			log.Ldate | log.Ltime | log.Lshortfile)
	}
}

func FileTreeStr(relativePath string, fileList FileList) string {
	sort.Sort(ByName(fileList))
	var output string
	for _, fileDetails := range fileList {
		filepath := relativePath + "/" + fileDetails.Name
		output += fmt.Sprintf("%32s %s % 14d %s\n",
			fileDetails.Md5sum, fileDetails.Id,
			fileDetails.Size, color.YellowString(filepath),
		)
		if len(fileDetails.Children) > 0 {
			output += FileTreeStr(relativePath + "/" + fileDetails.Name,
				fileDetails.Children)
		}
	}
	return output
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func CreatePath(path string) error {
	return os.MkdirAll(path, 0750)
}
