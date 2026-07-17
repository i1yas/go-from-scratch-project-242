package code

import (
	"code"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	type FileTestCase struct {
		filename string
		size     int64
	}

	cases := []FileTestCase{
		{filename: "test.txt", size: 6},
		{filename: "test2.txt", size: 11},
		{filename: "empty.txt", size: 0},
	}

	for _, tCase := range cases {
		t.Run(tCase.filename, func(t *testing.T) {
			path, err := filepath.Abs("../testdata/" + tCase.filename)
			if err != nil {
				t.Fatalf("File '%s' not found in testdata", tCase.filename)
			}

			size, err := code.GetPathSize(path, false)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	type DirTestCase struct {
		dirname       string
		size          int64
		includeHidden bool
	}

	cases := []DirTestCase{
		{dirname: "empty_dir", size: 0},
		{dirname: "dir_with_one_file", size: 6},
		{dirname: "dir_with_nested_dir", size: 6},
		{dirname: "dir_with_hidden_files", size: 6},
		{dirname: "dir_with_hidden_files", size: 18, includeHidden: true},
	}

	for _, tCase := range cases {
		t.Run(tCase.dirname, func(t *testing.T) {
			path, err := filepath.Abs("../testdata/" + tCase.dirname)
			if err != nil {
				t.Fatalf("File '%s' not found in testdata", tCase.dirname)
			}

			size, err := code.GetPathSize(path, tCase.includeHidden)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}
