package errors

import (
	"fmt"
)

type MusicFlacTagError struct {
	ext string
}

func (e *MusicFlacTagError) Error() string {
	return fmt.Sprintf("读取flac标签错误：%s", e.ext)
}

func NewMusicFlacTagError(ext string) *MusicFlacTagError {
	return &MusicFlacTagError{ext: ext}
}
