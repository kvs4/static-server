package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileExists(t *testing.T) {
	fs := NewFileService("./test_files")

	// Preparing text files
	err := os.Mkdir("./test_files", 0755)
	require.NoError(t, err, "Folder created error: %v", err)
	defer os.RemoveAll("./test_files")
	f, err := os.Create("./test_files/test.html")
	require.NoError(t, err, "File created error: %v", err)
	f.Close()

	// Checks
	if !fs.FileExists("test.html") {
		t.Errorf("Expected file to exist, but it does not.")
	}

	if fs.FileExists("missing.html") {
		t.Errorf("Expected file to not exist, but it does.")
	}
}
