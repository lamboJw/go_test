package media

import (
	"fmt"
	"go_test/music_player/errors"
	"go_test/music_player/utils"
	"path/filepath"
)

type Music struct {
	Player mediaInterface
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
	err = p.init()
	if err != nil {
		return nil, err
	}
	fmt.Printf("player: %+v\n", p)
	m := &Music{Player: p}
	return m, nil
}
