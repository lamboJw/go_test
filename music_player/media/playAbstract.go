package media

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"go_test/music_player/errors"
	"os"
	"time"
)

type playAbstract struct {
	mediaIniter
	//文件属性
	name string
	size int64
	//多媒体属性
	id       string
	title    string
	album    string
	artist   string
	year     string
	comment  string
	genre    uint8
	sort     int64
	filepath string
	fp       *os.File
	streamer beep.StreamSeekCloser
	format   beep.Format
	position int
}

func (p *playAbstract) Name() string {
	return p.name
}

func (p *playAbstract) Size() int64 {
	return p.size
}

func (p *playAbstract) Id() string {
	return p.id
}

func (p *playAbstract) Title() string {
	return p.title
}

func (p *playAbstract) Album() string {
	return p.album
}

func (p *playAbstract) Artist() string {
	return p.artist
}

func (p *playAbstract) Year() string {
	return p.year
}

func (p *playAbstract) Comment() string {
	return p.comment
}

func (p *playAbstract) Genre() uint8 {
	return p.genre
}

func (p *playAbstract) Sort() int64 {
	return p.sort
}

func (p *playAbstract) open() error {
	file, err := os.Open(p.filepath)
	if err != nil {
		return errors.NewMusicNotExistError(p.name)
	}
	p.fp = file
	return nil
}

func (p *playAbstract) firstPlay() error {
	err := speaker.Init(p.format.SampleRate, p.format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("初始化扬声器失败：", err)
		return err
	}
	return nil
}

func (p *playAbstract) init() error {
	var err error
	err = p.open()
	if err != nil {
		return err
	}
	err = p.initMediaInfo()
	if err != nil {
		return err
	}
	defer func() {
		err := p.fp.Close()
		if err != nil {
			fmt.Println("关闭文件", p.name, "失败：", err)
			return
		}
	}()
	return nil
}

func (p *playAbstract) doPlay(streamer beep.StreamSeekCloser) {
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (p *playAbstract) Pause() error {
	//TODO 记录当前position，然后speaker.Clear()
	p.position = p.streamer.Position()
	return nil
}

func (p *playAbstract) resume() error {
	//TODO 根据当前position，streamer.Seek(position)，然后p.doPlay(新streamer)
	return nil
}
