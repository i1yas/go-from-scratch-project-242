package internal

import "fmt"

type sizeUnitEntry struct {
	size    int64
	postfix string
}

var sizeUnits = []sizeUnitEntry{
	{size: int64(1) << 0, postfix: "B"},
	{size: int64(1) << 10, postfix: "KB"},
	{size: int64(1) << 20, postfix: "MB"},
	{size: int64(1) << 30, postfix: "GB"},
	{size: int64(1) << 40, postfix: "TB"},
	{size: int64(1) << 50, postfix: "PB"},
	{size: int64(1) << 60, postfix: "EB"},
}

func FormatSize(size int64, isHumanReadable bool) string {
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

func FormatCLIOutput(path string, size int64, isHumanReadable bool) string {
	return fmt.Sprintf("%s\t%s", path, FormatSize(size, isHumanReadable))
}
