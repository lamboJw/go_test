package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func SearchFiles(dir string) ([]string, error) {
	var files []string
	fsys := os.DirFS(dir)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		typeNum := ExtToMusicType(ext)
		if path != "." && typeNum != Unknown {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func TrimByteSlice(s []byte) []byte {
	index := len(s)
	zero := byte(0)
	for k, v := range s {
		if v == zero {
			index = k
			break
		}
	}
	fmt.Println(s, index, s[:index])
	return s[:index]
}
