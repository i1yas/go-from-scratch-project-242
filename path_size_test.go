package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	type FileTestCase struct {
		filename string
		size     string
	}

	cases := []FileTestCase{
		{filename: "test.txt", size: "6B"},
		{filename: "test2.txt", size: "11B"},
		{filename: "empty.txt", size: "0B"},
	}

	for _, tCase := range cases {
		t.Run(tCase.filename, func(t *testing.T) {
			path, err := filepath.Abs(filepath.Join("testdata", tCase.filename))
			if err != nil {
				t.Fatalf("File '%s' not found in testdata", tCase.filename)
			}

			size, err := GetPathSize(path, false, false, false)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}

func TestGetPathSize_Dir(t *testing.T) {
	type DirTestCase struct {
		dirname       string
		size          string
		includeHidden bool
	}

	cases := []DirTestCase{
		{dirname: "empty_dir", size: "0B"},
		{dirname: "dir_with_one_file", size: "6B"},
		{dirname: "dir_with_nested_dir", size: "6B"},
		{dirname: "dir_with_hidden_files", size: "6B"},
		{dirname: "dir_with_hidden_files", size: "18B", includeHidden: true},
	}

	for _, tCase := range cases {
		t.Run(tCase.dirname, func(t *testing.T) {
			path, err := filepath.Abs(filepath.Join("testdata", tCase.dirname))
			if err != nil {
				t.Fatalf("Directory '%s' not found in testdata", tCase.dirname)
			}

			size, err := GetPathSize(path, false, false, tCase.includeHidden)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}

func TestGetPathSize_Dir_Recursive(t *testing.T) {
	type DirTestCase struct {
		dirname       string
		size          string
		includeHidden bool
	}

	cases := []DirTestCase{
		{dirname: "empty_dir", size: "0B"},
		{dirname: "dir_with_one_file", size: "6B"},
		{dirname: "dir_with_nested_dir", size: "12B"},
		{dirname: "dir_with_hidden_files", size: "6B"},
		{dirname: "dir_with_hidden_files", size: "18B", includeHidden: true},
	}

	for _, tCase := range cases {
		t.Run(tCase.dirname, func(t *testing.T) {
			path, err := filepath.Abs(filepath.Join("testdata", tCase.dirname))
			if err != nil {
				t.Fatalf("Directory '%s' not found in testdata", tCase.dirname)
			}

			size, err := GetPathSize(path, true, false, tCase.includeHidden)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}
