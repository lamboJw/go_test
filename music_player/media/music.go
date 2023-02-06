package media

import (
	"go_test/music_player/errors"
	"go_test/music_player/interfaces"
	"go_test/music_player/types"
	"path/filepath"
)

var mediaTypeMap = make(map[string]mediaCreator)

type mediaCreator func(path string) interfaces.MediaInterface

/*
注册工厂方法
*/
func mediaRegister(typeStr string, factory mediaCreator) {
	mediaTypeMap[typeStr] = factory
}

/*
根据类型调用对应的工厂方法
*/
func createMedia(typeStr string, path string) (interfaces.MediaInterface, error) {
	if factory, ok := mediaTypeMap[typeStr]; ok {
		return factory(path), nil
	} else {
		return nil, errors.NewMusicTypeError(typeStr)
	}
}

type Music struct {
	player interfaces.MediaInterface
}

func NewMusic(path string) (*Music, error) {
	ext := filepath.Ext(path)
	var p interfaces.MediaInterface
	var err error
	p, err = createMedia(types.ExtToMusicType(ext).String(), path)
	if err != nil {
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
