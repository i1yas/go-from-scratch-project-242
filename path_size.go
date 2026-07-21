package code

import (
	"fmt"
	"os"
	"strings"
)

func GetPathSize(path string, isRecursive bool, isHumanReadable bool, includeHidden bool) (string, error) {
	size, err := getPathSizeRaw(path, isRecursive, includeHidden)
	if err != nil {
		return "", err
	}
	return formatFileSize(size, isHumanReadable), nil
}

func getPathSizeRaw(path string, isRecursive bool, includeHidden bool) (int64, error) {
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

			dirSize, err := getPathSizeRaw(path+"/"+entry.Name(), isRecursive, includeHidden)
			if err != nil {
				return 0, err
			}
			totalSize += dirSize

			continue
		}

		totalSize += fileInfo.Size()

	}

	return totalSize, nil
}

type SizeUnitEntry struct {
	size    int64
	postfix string
}

var sizeUnits = []SizeUnitEntry{
	{size: int64(1) << 0, postfix: "B"},
	{size: int64(1) << 10, postfix: "KB"},
	{size: int64(1) << 20, postfix: "MB"},
	{size: int64(1) << 30, postfix: "GB"},
	{size: int64(1) << 40, postfix: "TB"},
	{size: int64(1) << 50, postfix: "PB"},
	{size: int64(1) << 60, postfix: "EB"},
}

func formatFileSize(size int64, isHumanReadable bool) string {
	if !isHumanReadable {
		return fmt.Sprintf("%dB", size)
	}

	unit := sizeUnits[0]

	for _, u := range sizeUnits {
		if size < u.size {
			break
		}
		unit = u
	}

	sizeInUnit := float32(size) / float32(unit.size)

	if unit.postfix == "B" {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", sizeInUnit, unit.postfix)
}
