package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, isRecursive bool, includeHidden bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("Failed to read info for path '%s': %w", path, err)
	}

	if !fileInfo.IsDir() {
		return fileInfo.Size(), nil
	}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("Failed to read dir '%s': %w", path, err)
	}

	totalSize := int64(0)
	for _, entry := range dirEntries {
		fileInfo, err := entry.Info()
		if err != nil {
			dirEntryPath := filepath.Join(path, entry.Name())
			return 0, fmt.Errorf("Failed to read dir entry '%s': %w", dirEntryPath, err)
		}

		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if fileInfo.IsDir() {
			if !isRecursive {
				continue
			}

			dirSize, err := GetPathSize(path+"/"+entry.Name(), isRecursive, includeHidden)
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
