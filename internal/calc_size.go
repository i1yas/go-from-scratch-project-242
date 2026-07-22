package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, isRecursive bool, includeHidden bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read info for path '%s': %w", path, err)
	}

	fileMode := fileInfo.Mode()

	if isSupportedFileType(fileMode) {
		return fileInfo.Size(), nil
	}

	if !fileInfo.IsDir() {
		return 0, fmt.Errorf("got unsupported file type: %s", getHumanReadableFileType(fileMode))
	}

	totalSize := int64(0)
	err = filepath.WalkDir(path, func(entryPath string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entryPath == path {
			return nil
		}

		fileInfo, err := entry.Info()
		if err != nil {
			return fmt.Errorf("failed to read dir entry '%s': %w", entryPath, err)
		}

		if fileInfo.IsDir() {
			if !isRecursive {
				return filepath.SkipDir
			}

			if isHidden(entry.Name()) && !includeHidden {
				return filepath.SkipDir
			}

			return nil
		}

		if isHidden(entry.Name()) && !includeHidden {
			return nil
		}

		fileMode := fileInfo.Mode()

		if isSupportedFileType(fileMode) {
			totalSize += fileInfo.Size()
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to walk dir '%s': %w", path, err)
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
