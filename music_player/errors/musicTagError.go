package errors

import "fmt"

type musicTagError struct {
	filename string
}

func (e *musicTagError) Error() string {
	return fmt.Sprintf("%s文件Tag标记错误", e.filename)
}

func NewMusicTagError(filename string) *musicTagError {
	return &musicTagError{filename: filename}
}
