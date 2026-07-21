package code

import (
	"code/internal"
)

func GetPathSize(path string, isRecursive bool, isHumanReadable bool, includeHidden bool) (string, error) {
	size, err := internal.GetPathSize(path, isRecursive, includeHidden)
	if err != nil {
		return "", err
	}

	return internal.FormatSize(size, isHumanReadable), nil
}
