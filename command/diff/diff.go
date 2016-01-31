package diff

import (
	. "github.com/taryk/gdtool/core"
)

func deepCopy(fileList FileList) FileList {
	var fileListCopy FileList
	for _, fileDetails := range fileList {
		fileListCopy = append(
			fileListCopy,
			&FileDetails{
				Id: fileDetails.Id,
				Name: fileDetails.Name,
				Path: fileDetails.Path,
				Size: fileDetails.Size,
				Md5sum: fileDetails.Md5sum,
				IsDir: fileDetails.IsDir,
				Children: deepCopy(fileDetails.Children),
			},
		)
	}
	return fileListCopy
}

func Compare(remoteFiles, localFiles FileList) (FileList, FileList) {
	var uniqueRemoteFiles FileList
	var uniqueLocalFiles FileList
	remoteFiles = uniqueFileList(deepCopy(remoteFiles))
	localFiles = uniqueFileList(deepCopy(localFiles))
	foundLocalFiles := make(map[int]interface{})
	for _, remoteFile := range remoteFiles {
		foundTheSameLocalFile := false
		if len(localFiles) > 0 {
			LocalFile:
			for i, localFile := range localFiles {
				// Skip local files which are already found.
				if _, ok := foundLocalFiles[i]; ok { continue LocalFile }
				if remoteFile.Name != localFile.Name { continue LocalFile }
				if remoteFile.Md5sum != localFile.Md5sum {
					Warn.Printf("Md5sums don't match: %+v || %+v.",
						remoteFile, localFile)
					break LocalFile
				}
				// OK, if it's not a folder, consider the files to be the same.
				if !remoteFile.IsDir {
					foundTheSameLocalFile = true
					foundLocalFiles[i] = nil
					break LocalFile
				}
				// If they are folders, let's compare them.
				remoteFile.Children, localFile.Children = Compare(
					remoteFile.Children,
					localFile.Children,
				)
				// If there are no unique remote nested files, consider the
				// remote folder to be the same as a local one.
				if len(remoteFile.Children) == 0 {
					foundTheSameLocalFile = true
				}
				// If there are no unique local nested files, ignore the folder
				// next time we go through the list of local files.
				if len(localFile.Children) == 0 {
					foundLocalFiles[i] = nil
				}
				break LocalFile
			}
		}
		if !foundTheSameLocalFile {
			uniqueRemoteFiles = append(uniqueRemoteFiles, remoteFile)
		}
	}
	// Check if there are any unique local files.
	for i, localFile := range localFiles {
		if _, ok := foundLocalFiles[i]; !ok {
			uniqueLocalFiles = append(uniqueLocalFiles, localFile)
		}
	}
	return uniqueRemoteFiles, uniqueLocalFiles
}

func uniqueFileList(inputFileList FileList) FileList {
	var result FileList
	seen := map[string]string{}
	for _, fileDetails := range inputFileList {
		if _, Md5sum := seen[fileDetails.Name]; !Md5sum {
			result = append(result, fileDetails)
			seen[fileDetails.Name] = fileDetails.Md5sum
		} else {
			Warn.Printf("Duplicate %+v.", fileDetails)
		}
	}
	return result
}
