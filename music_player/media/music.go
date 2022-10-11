package media

import (
	"go_test/music_player/errors"
	"go_test/music_player/utils"
	"path/filepath"
)

type Music struct {
	Player mediaInterface
}

func (m *Music) Name() string {
	return m.Player.Name()
}

func (m *Music) Size() int64 {
	return m.Player.Size()
}

func (m *Music) Id() string {
	return m.Player.Id()
}

func (m *Music) Title() string {
	return m.Player.Title()
}

func (m *Music) Artist() string {
	return m.Player.Artist()
}

func (m *Music) Album() string {
	return m.Player.Album()
}

func (m *Music) Year() string {
	return m.Player.Year()
}

func (m *Music) Genre() string {
	return m.Player.Genre()
}

func (m *Music) Sort() int64 {
	return m.Player.Sort()
}

func (m *Music) Index() int64 {
	return m.Player.Index()
}

func NewMusic(path string) (*Music, error) {
	ext := filepath.Ext(path)
	var p mediaInterface
	var err error
	typeNum := utils.ExtToMusicType(ext)
	switch typeNum {
	case utils.Mp3:
		p = newMp3Player(path)
	case utils.Wav:
		p = newWavPlayer(path)
	default:
		err = errors.NewMusicTypeError(ext)
		return nil, err
	}
	m := &Music{Player: p}
	err = m.init()
	if err != nil {
		return nil, err
	}
	//fmt.Printf("player: %+v\n", p)
	return m, nil
}

func (m *Music) init() error {
	var err error
	_, err = m.Player.Fp()
	if err != nil {
		return err
	}
	defer m.Player.CloseFp()
	err = m.Player.initMediaInfo()
	if err != nil {
		return err
	}
	return nil
}

func (m *Music) Play() error {
	var err error
	_, err = m.Player.Streamer()
	if err != nil {
		return err
	}
	defer m.Player.CloseStreamer()
	err = m.Player.Play()
	if err != nil {
		return err
	}
	return nil
}

func (m *Music) Pause() error {
	var err error
	err = m.Player.Pause()
	if err != nil {
		return err
	}
	return nil
}
