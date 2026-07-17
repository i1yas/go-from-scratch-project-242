package code

import (
	"os"
	"strings"
)

func GetPathSize(path string, includeHidden bool) (int64, error) {
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

		if fileInfo.IsDir() {
			continue
		}

		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		totalSize += fileInfo.Size()

	}

	return totalSize, nil
}
