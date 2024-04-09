package utils

import (
	"io/fs"
	"path/filepath"
)

func GetFiles(dir string, ext string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
