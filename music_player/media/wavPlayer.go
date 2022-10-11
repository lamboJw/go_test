package media

import (
	"github.com/faiface/beep"
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

func (f *wavPlayer) Streamer() (beep.StreamSeekCloser, error) {
	if f.streamer != nil {
		return f.streamer, nil
	}
	err := f.initStreamer()
	if err != nil {
		return nil, err
	}
	return f.streamer, nil
}

func (f *wavPlayer) initMediaInfo() error {
	return nil
}
