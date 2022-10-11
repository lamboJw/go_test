package media

import (
	"github.com/faiface/beep"
	"os"
)

type MediaInfoGetter interface {
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

type mediaInfoSetter interface {
	SetIndex(index int64)
}

type mediaFileGetSetter interface {
	Fp() (*os.File, error)
	CloseFp()
	Streamer() (beep.StreamSeekCloser, error)
	CloseStreamer()
}
