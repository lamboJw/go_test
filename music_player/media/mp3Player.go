package media

import (
	"crypto/md5"
	"fmt"
	"github.com/faiface/beep/mp3"
	"go_test/music_player/errors"
	"strings"
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
	if f.streamer != nil {
		return nil
	}
	err := f.open()
	if err != nil {
		return err
	}
	defer func() {
		err := f.fp.Close()
		if err != nil {
			fmt.Println("关闭文件", f.name, "失败：", err)
			return
		}
	}()
	streamer, format, err := mp3.Decode(f.fp)
	if err != nil {
		return err
	}
	f.streamer = streamer
	f.format = format
	return nil
}

func (f *mp3Player) initMediaInfo() error {
	fInfo, _ := f.fp.Stat()
	f.sort = fInfo.ModTime().Unix()
	f.name = fInfo.Name()
	f.size = fInfo.Size()
	err := f.id3V2Info()
	if err != nil {
		return err
	}
	return nil
}

func (f *mp3Player) id3V1Info() error {
	id3 := make([]byte, 128)
	_, err := f.fp.ReadAt(id3, f.size-128)
	if err != nil {
		return err
	}
	if string(id3[:3]) != "TAG" {
		return errors.NewMusicTagError(f.name)
	}
	f.title = strings.Trim(string(id3[3:33]), "\x00")
	f.album = strings.Trim(string(id3[33:63]), "\x00")
	f.artist = strings.Trim(string(id3[63:93]), "\x00")
	f.year = string(id3[93:97])
	f.comment = strings.Trim(string(id3[97:127]), "\x00")
	f.genre = id3[127]
	f.id = fmt.Sprintf("%x", md5.Sum([]byte(f.title+f.album+f.artist+f.year+f.comment)))
	return nil
}

func (f *mp3Player) id3V2Info() error {
	signHeader := make([]byte, 10)
	_, err := f.fp.Read(signHeader)
	if err != nil {
		return err
	}
	if string(signHeader[:3]) != "ID3" {
		return errors.NewMusicTagError(f.name)
	}
	version := string(signHeader[3:4])
	fmt.Println(version)
	return nil
}

/*func (f *mp3Player) init() error {
	var err error
	err = f.open()
	if err != nil {
		return err
	}
	err = f.initMediaInfo()
	if err != nil {
		return err
	}
	defer func() {
		err := f.fp.Close()
		if err != nil {
			fmt.Println("关闭文件", f.name, "失败：", err)
			return
		}
	}()
	return nil
}*/

func (f *mp3Player) Play() error {
	var err error
	defer func() {
		err := f.streamer.Close()
		if err != nil {
			fmt.Println("关闭流媒体失败：", err)
			return
		}
	}()
	if f.position == 0 {
		err = f.firstPlay()
	} else {
		err = f.resume()
	}
	if err != nil {
		return nil
	}
	f.doPlay(f.streamer)
	return nil
}
