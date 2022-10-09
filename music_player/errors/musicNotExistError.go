package errors

import "fmt"

type MusicNotExistError struct {
	filename string
}

func (e *MusicNotExistError) Error() string {
	return fmt.Sprintf("找不到%s文件", e.filename)
}

func NewMusicNotExistError(filename string) *MusicNotExistError {
	return &MusicNotExistError{filename: filename}
}
