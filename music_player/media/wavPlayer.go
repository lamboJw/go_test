package media

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"go_test/music_player/interfaces"
	"go_test/music_player/types"
)

type wavPlayer struct {
	*basePlayer
}

func init() {
	mediaRegister(types.Wav.String(), newWavPlayer)
}

func newWavPlayer(filepath string) interfaces.MediaInterface {
	instance := &wavPlayer{
		basePlayer: &basePlayer{
			filepath: filepath,
		},
	}
	return instance
}

func (f *wavPlayer) InitStreamer() error {
	streamer, format, err := wav.Decode(f.fp)
	if err != nil {
		return err
	}
	f.streamer = streamer
	f.format = format
	return nil
}

func (f *wavPlayer) Streamer() (beep.StreamSeekCloser, error) {
	if f.streamer != nil {
		return f.streamer, nil
	}
	err := f.InitStreamer()
	if err != nil {
		return nil, err
	}
	return f.streamer, nil
}

func (f *wavPlayer) InitMediaInfo() error {
	return nil
}
