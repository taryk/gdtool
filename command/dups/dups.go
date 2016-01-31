package dups

import (
	"fmt"
	"strconv"
	"sort"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"

	. "github.com/taryk/gdtool/core"
)

func PrintDuplicates(fileTreeMap FileTreeMap) {
	canBeDeleted := FindDuplicates("", fileTreeMap)
	var bytesToRelease uint64
	color.Yellow("Files to delete:\n")
	for _, fileDetailsMap := range canBeDeleted {
		filename := fileDetailsMap.Path + "/" + fileDetailsMap.Name
		fmt.Printf("%32s % 14d %s %s\n",
			fileDetailsMap.Md5sum, fileDetailsMap.Size,
			color.YellowString(filename),
			fileDetailsMap.Id,
		)
		bytesToRelease += fileDetailsMap.Size
	}
	fmt.Printf("There are %s files to delete. %s can be released\n",
		color.YellowString(strconv.Itoa(len(canBeDeleted))),
		color.YellowString(humanize.Bytes(bytesToRelease)),
	)
}

func FindDuplicates(relativePath string,
	fileTreeMap FileTreeMap) []*FileDetailsMap {

	var filenames []string
	for filename := range fileTreeMap {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)
	var canBeDeleted []*FileDetailsMap
	filename:
	for _, filename := range filenames {
		if len(fileTreeMap[filename]) == 1 {
			fileDetailsMap := fileTreeMap[filename][0]
			if fileDetailsMap.IsDir {
				path := ""
				if relativePath != "/" {
					path = relativePath
				}
				path += "/" + fileDetailsMap.Name
				canBeDeleted = append(canBeDeleted,
					FindDuplicates(path, fileDetailsMap.Children)...)
			}
			continue filename
		}

		color.Yellow(relativePath + "/" + filename)
		md5map := make(map[string][]*FileDetailsMap)
		for _, fileDetailsMap := range fileTreeMap[filename] {
			md5map[fileDetailsMap.Md5sum] = append(
				md5map[fileDetailsMap.Md5sum],
				fileDetailsMap,
			)
		}

		for md5sum := range md5map {
			fmt.Printf("%32s " +
				color.YellowString("(count: %d)", len(md5map[md5sum])) + "\n",
				md5sum,
			)
			for i, fileDetailsMap := range md5map[md5sum] {
				fmt.Printf(" % 29s %d",
					color.YellowString(fileDetailsMap.Id),
					fileDetailsMap.Size,
				)
				if fileDetailsMap.IsDir {
					fmt.Print(" ", len(fileDetailsMap.Children))
				}
				fmt.Println()
				if i > 0 {
					fileDetailsMap.Path = relativePath
					canBeDeleted = append(canBeDeleted, fileDetailsMap)
				}
			}
		}

		fmt.Println()
	}

	return canBeDeleted
}

func GroupByName(inputFileList FileList) FileTreeMap {
	names := make(FileTreeMap)
	for _, fileDetails := range inputFileList {
		var fileDetailsMap FileDetailsMap

		fileDetailsMap.Id = fileDetails.Id
		fileDetailsMap.Name = fileDetails.Name
		fileDetailsMap.IsDir = fileDetails.IsDir
		fileDetailsMap.Md5sum = fileDetails.Md5sum
		fileDetailsMap.Path = fileDetails.Path
		fileDetailsMap.Size = fileDetails.Size

		if fileDetails.IsDir {
			fileDetailsMap.Children = GroupByName(fileDetails.Children)
		}
		names[fileDetails.Name] = append(names[fileDetails.Name],
			&fileDetailsMap)
	}
	return names
}
