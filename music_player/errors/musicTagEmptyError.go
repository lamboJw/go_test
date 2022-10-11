package errors

import "fmt"

type MusicTagEmptyError struct {
	filename string
	offset   int64
}

func (e *MusicTagEmptyError) Error() string {
	return fmt.Sprintf("%s文件Tag标记为空，当前偏移量：%d", e.filename, e.offset)
}

func NewMusicTagEmptyError(filename string, offset int64) *MusicTagEmptyError {
	return &MusicTagEmptyError{filename: filename, offset: offset}
}
