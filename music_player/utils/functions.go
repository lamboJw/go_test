package utils

import (
	"fmt"
	"go_test/music_player/types"
	"io/fs"
	"os"
	"path/filepath"
	"unicode"
)

func SearchFiles(dir string) ([]string, error) {
	var files []string
	fsys := os.DirFS(dir)
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		typeNum := types.ExtToMusicType(ext)
		if path != "." && typeNum != types.Unknown {
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

func GetOffsetBit(b byte, offset uint8) uint8 {
	d := (b << (8 - offset)) >> 7
	return d
}

func SetOffsetBit(b byte, offset uint8, setNum uint8) byte {
	var d byte
	var one byte = 1
	if setNum == 1 {
		d = b | (one << (offset - 1))
	} else {
		d = b & ^(one << (offset - 1))
	}
	return d
}

func ChineseCount(str1 string) (count int) {
	for _, char := range str1 {
		if unicode.Is(unicode.Han, char) {
			count++
		}
	}
	return
}
