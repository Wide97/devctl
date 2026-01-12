package fileutil

import (
	"os"
)

// ListFiles restituisce i nomi dei file in una directory
func ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}
	return names, nil
}
