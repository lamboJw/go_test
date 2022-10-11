package media

import (
	"go_test/music_player/errors"
	"go_test/music_player/interfaces"
	"go_test/music_player/types"
	"path/filepath"
)

type Music struct {
	player interfaces.MediaInterface
}

func NewMusic(path string) (*Music, error) {
	ext := filepath.Ext(path)
	var p interfaces.MediaInterface
	var err error
	typeNum := types.ExtToMusicType(ext)
	switch typeNum {
	case types.Mp3:
		p = newMp3Player(path)
	case types.Wav:
		p = newWavPlayer(path)
	default:
		err = errors.NewMusicTypeError(ext)
		return nil, err
	}
	m := &Music{player: p}
	err = m.init()
	if err != nil {
		return nil, err
	}
	//fmt.Printf("player: %+v\n", p)
	return m, nil
}

func (m *Music) init() error {
	var err error
	_, err = m.player.Fp()
	if err != nil {
		return err
	}
	defer m.player.CloseFp()
	err = m.player.InitMediaInfo()
	if err != nil {
		return err
	}
	return nil
}

func (m *Music) Name() string {
	return m.player.Name()
}

func (m *Music) Size() int64 {
	return m.player.Size()
}

func (m *Music) Id() string {
	return m.player.Id()
}

func (m *Music) Title() string {
	return m.player.Title()
}

func (m *Music) Artist() string {
	return m.player.Artist()
}

func (m *Music) Album() string {
	return m.player.Album()
}

func (m *Music) Year() string {
	return m.player.Year()
}

func (m *Music) Genre() string {
	return m.player.Genre()
}

func (m *Music) Sort() int64 {
	return m.player.Sort()
}

func (m *Music) Index() int64 {
	return m.player.Index()
}

func (m *Music) Play() error {
	var err error
	_, err = m.player.Fp()
	if err != nil {
		return err
	}
	defer m.player.CloseFp()
	_, err = m.player.Streamer()
	if err != nil {
		return err
	}
	defer m.player.CloseStreamer()
	err = m.player.Play()
	if err != nil {
		return err
	}
	return nil
}

func (m *Music) Pause() error {
	var err error
	err = m.player.Pause()
	if err != nil {
		return err
	}
	return nil
}
