package media

import (
	"crypto/md5"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

type mp3Player struct {
	playAbstract
}

func newMp3Player(filepath string) mediaInterface {
	instance := &mp3Player{
		playAbstract: playAbstract{
			filepath: filepath,
		},
	}
	return instance
}

func (f *mp3Player) initStreamer() error {
	streamer, format, err := mp3.Decode(f.fp)
	if err != nil {
		return err
	}
	f.streamer = streamer
	f.format = format
	return nil
}

func (f *mp3Player) Streamer() (beep.StreamSeekCloser, error) {
	if f.streamer != nil {
		return f.streamer, nil
	}
	err := f.initStreamer()
	if err != nil {
		return nil, err
	}
	return f.streamer, nil
}

func (f *mp3Player) initMediaInfo() error {
	fInfo, _ := f.fp.Stat()
	f.sort = fInfo.ModTime().Unix()
	f.name = fInfo.Name()
	f.size = fInfo.Size()
	id3, err := newID3(f.fp)
	if err != nil {
		return err
	}
	f.title = id3.Title()
	f.artist = id3.Artist()
	f.album = id3.Album()
	f.year = id3.Album()
	f.genre = id3.Genre()
	f.id = fmt.Sprintf("%x", md5.Sum([]byte(f.title+f.artist+f.album+f.year+f.genre)))
	return nil
}
