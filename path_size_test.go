package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	type FileTestCase struct {
		name            string
		filename        string
		size            string
		isHumanReadable bool
		hasError        bool
	}

	cases := []FileTestCase{
		{name: "text.txt raw format", filename: "test.txt", size: "6B", isHumanReadable: false},
		{name: "text.txt human-readable", filename: "test.txt", size: "6B", isHumanReadable: true},

		{name: "test2.txt raw format", filename: "test2.txt", size: "11B", isHumanReadable: false},
		{name: "test2.txt human-readable", filename: "test2.txt", size: "11B", isHumanReadable: true},

		{name: "test3.txt raw format", filename: "test3.txt", size: "1343B", isHumanReadable: false},
		{name: "test3.txt human-readable", filename: "test3.txt", size: "1.3KB", isHumanReadable: true},

		{name: "empty.txt raw format", filename: "empty.txt", size: "0B", isHumanReadable: false},
		{name: "empty.txt human-readable", filename: "empty.txt", size: "0B", isHumanReadable: true},

		{name: "path does not exist", filename: "does_not_exist.txt", hasError: true},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			path, err := filepath.Abs(filepath.Join("testdata", tCase.filename))
			if err != nil {
				t.Fatalf("File '%s' not found in testdata", tCase.filename)
			}

			size, err := GetPathSize(path, false, tCase.isHumanReadable, false)

			if tCase.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tCase.size, size)
			}
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

func TestGetPathSize_Empty_Dir(t *testing.T) {

	type EmptyDirTestCase struct {
		name            string
		isHumanReadable bool
		includeHidden   bool
		size            string
	}

	cases := []EmptyDirTestCase{
		{name: "raw format", isHumanReadable: false, includeHidden: false, size: "0B"},
		{name: "human-readable format", isHumanReadable: true, includeHidden: false, size: "0B"},
		{name: "raw format and include hidden", isHumanReadable: false, includeHidden: true, size: "0B"},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			tempDir := t.TempDir()

			size, err := GetPathSize(tempDir, true, tCase.isHumanReadable, tCase.includeHidden)

			require.NoError(t, err)
			require.Equal(t, tCase.size, size)
		})
	}
}
