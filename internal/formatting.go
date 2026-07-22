package internal

import (
	"fmt"
)

type sizeUnitEntry struct {
	scale  int
	suffix string
}

func (u sizeUnitEntry) size() int64 {
	return int64(1) << (u.scale * 10)
}

var orderedSizeUnits = []sizeUnitEntry{
	{scale: 0, suffix: "B"},
	{scale: 1, suffix: "KB"},
	{scale: 2, suffix: "MB"},
	{scale: 3, suffix: "GB"},
	{scale: 4, suffix: "TB"},
	{scale: 5, suffix: "PB"},
	{scale: 6, suffix: "EB"},
}

func FormatSize(size int64, isHumanReadable bool) string {
	unit := orderedSizeUnits[0]

	if isHumanReadable {
		unit = pickUnit(size)
	}

	if unit.scale == 0 {
		return fmt.Sprintf("%d%s", size, unit.suffix)
	}

	sizeInUnit := float32(size) / float32(unit.size())
	return fmt.Sprintf("%.1f%s", sizeInUnit, unit.suffix)
}

func FormatCLIOutput(path string, size int64, isHumanReadable bool) string {
	return fmt.Sprintf("%s\t%s", FormatSize(size, isHumanReadable), path)
}

func pickUnit(size int64) sizeUnitEntry {
	unit := orderedSizeUnits[0]
	for _, u := range orderedSizeUnits[1:] {
		if size < u.size() {
			break
		}
		unit = u
	}
	return unit
}
