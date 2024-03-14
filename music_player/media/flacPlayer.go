package media

import (
	"crypto/md5"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"go_test/music_player/interfaces"
	"go_test/music_player/types"
)

type flacPlayer struct {
	*basePlayer
}

func init() {
	mediaRegister(types.Flac.String(), newFlacPlayer)
}

func newFlacPlayer(filepath string) interfaces.MediaInterface {
	instance := &flacPlayer{
		basePlayer: &basePlayer{
			filepath: filepath,
		},
	}
	return instance
}

func (f *flacPlayer) InitStreamer() error {
	streamer, format, err := flac.Decode(f.fp)
	if err != nil {
		return err
	}
	f.streamer = streamer
	f.format = format
	return nil
}

func (f *flacPlayer) Streamer() (beep.StreamSeekCloser, error) {
	if f.streamer != nil {
		return f.streamer, nil
	}
	err := f.InitStreamer()
	if err != nil {
		return nil, err
	}
	return f.streamer, nil
}

func (f *flacPlayer) InitMediaInfo() error {
	fInfo, _ := f.fp.Stat()
	f.sort = fInfo.ModTime().Unix()
	f.name = fInfo.Name()
	f.size = fInfo.Size()
	metaData, err := newFlacMetaData(f.fp)
	if err != nil {
		return err
	}
	f.title = metaData.Title()
	f.artist = metaData.Artist()
	f.album = metaData.Album()
	f.year = metaData.Year()
	f.genre = metaData.Genre()
	f.id = fmt.Sprintf("%x", md5.Sum([]byte(f.title+f.artist+f.album+f.year+f.genre)))
	return nil
}
