package fileutil

import (
	"os"
	"testing"
)

func TestFileList(t *testing.T) {
	// Temporary directory for testing
	dir := t.TempDir()

	// Create example files
	fileNames := []string{"a.txt", "b.txt", "c.txt"}
	for _, name := range fileNames {
		path := dir + "/" + name
		if err := os.WriteFile(path, []byte("test"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}
	}

	// Call to the function to test
	files, err := ListFiles(dir)
	if err != nil {
		t.Fatalf("ListFiles returned error: %v", err)
	}

	// Check file list
	for _, name := range fileNames {
		found := false
		for _, f := range files {
			if f == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected file %s not found in result", name)
		}
	}

	// Test empty folder
	emptyDir := t.TempDir()
	files, err = ListFiles(emptyDir)
	if err != nil {
		t.Fatalf("ListFiles returned error on empty dir: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}
}
