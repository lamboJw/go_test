package media

import (
	"github.com/faiface/beep/wav"
)

type wavPlayer struct {
	playAbstract
}

func newWavPlayer(filepath string) *wavPlayer {
	instance := &wavPlayer{
		playAbstract: playAbstract{
			filepath: filepath,
		},
	}
	return instance
}

func (f *wavPlayer) initStreamer() error {
	streamer, format, err := wav.Decode(f.fp)
	if err != nil {
		return err
	}
	f.streamer = streamer
	f.format = format
	return nil
}

func (f *wavPlayer) Play() error {
	return nil
}
