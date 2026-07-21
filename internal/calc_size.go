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

	fileMode := fileInfo.Mode()

	if isSupportedFileType(fileMode) {
		return fileInfo.Size(), nil
	}

	if !fileInfo.IsDir() {
		return 0, fmt.Errorf("Got unsupported file type: %s", getHumanReadableFileType(fileMode))
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

		if !includeHidden && isHidden(entry.Name()) {
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

		fileMode := fileInfo.Mode()

		if isSupportedFileType(fileMode) {
			totalSize += fileInfo.Size()
		}
	}

	return totalSize, nil
}

func getHumanReadableFileType(mode os.FileMode) string {
	fileType := mode.Type()
	switch {
	case mode.IsRegular():
		return "regular file"
	case fileType == os.ModeDir:
		return "directory"
	case fileType == os.ModeSymlink:
		return "symbolic link"
	case fileType == os.ModeDevice:
		return "device"
	case fileType == os.ModeNamedPipe:
		return "named pipe"
	case fileType == os.ModeSocket:
		return "socket"
	default:
		return "unknown type"
	}
}

func isSupportedFileType(fileMode os.FileMode) bool {
	return fileMode.IsRegular() || fileMode.Type() == os.ModeSymlink
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}
