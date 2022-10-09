package errors

import (
	"fmt"
)

type MusicTypeError struct {
	ext string
}

func (e *MusicTypeError) Error() string {
	return fmt.Sprintf("未知的音乐类型：%s", e.ext)
}

func NewMusicTypeError(ext string) *MusicTypeError {
	return &MusicTypeError{ext: ext}
}
