package code

import (
	"os"
	"strings"
)

func GetPathSize(path string, includeHidden bool, isRecursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	totalSize := int64(0)
	for _, entry := range dirEntries {
		fileInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if fileInfo.IsDir() {
			if !isRecursive {
				continue
			}

			dirSize, err := GetPathSize(path+"/"+entry.Name(), includeHidden, isRecursive)
			if err != nil {
				return 0, nil
			}
			totalSize += dirSize

			continue
		}

		totalSize += fileInfo.Size()

	}

	return totalSize, nil
}
