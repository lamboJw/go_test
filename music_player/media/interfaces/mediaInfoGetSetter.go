package interfaces

import (
	"github.com/faiface/beep"
	"os"
)

type InfoGetter interface {
	Name() string
	Size() int64
	Id() string
	Title() string
	Artist() string
	Album() string
	Year() string
	Genre() string
	Sort() int64
	Index() int64
}

type InfoSetter interface {
	SetIndex(index int64)
}

type FileGetSetter interface {
	Fp() (*os.File, error)
	CloseFp()
	Streamer() (beep.StreamSeekCloser, error)
	CloseStreamer()
}
