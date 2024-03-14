package interfaces

import (
	"github.com/faiface/beep"
	"os"
)

type MediaInterface interface {
	MediaPlayer
	MediaInfoGetter
	MediaInfoSetter
	MediaFileGetSetter
	MediaIniter
}

type MediaInfoGetterAndPlayer interface {
	MediaPlayer
	MediaInfoGetter
}

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
	Playing() bool
}

type MediaInfoSetter interface {
	SetIndex(index int64)
}

type MediaFileGetSetter interface {
	Fp() (*os.File, error)
	CloseFp()
	Streamer() (beep.StreamSeekCloser, error)
	CloseStreamer()
}

type MediaIniter interface {
	InitStreamer() error
	InitMediaInfo() error
}

type MediaPlayer interface {
	Play() error
	Pause() error
}
