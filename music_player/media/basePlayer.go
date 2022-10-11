package media

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"go_test/music_player/errors"
	"os"
	"time"
)

type basePlayer struct {
	//文件属性
	name string
	size int64
	//多媒体属性
	id     string
	title  string
	album  string
	artist string
	year   string
	genre  string
	//播放器属性
	sort     int64
	index    int64
	filepath string
	fp       *os.File
	streamer beep.StreamSeekCloser
	format   beep.Format
	position int
}

func (p *basePlayer) SetIndex(index int64) {
	p.index = index
}

func (p *basePlayer) Index() int64 {
	return p.index
}

func (p *basePlayer) Name() string {
	return p.name
}

func (p *basePlayer) Size() int64 {
	return p.size
}

func (p *basePlayer) Id() string {
	return p.id
}

func (p *basePlayer) Title() string {
	return p.title
}

func (p *basePlayer) Album() string {
	return p.album
}

func (p *basePlayer) Artist() string {
	return p.artist
}

func (p *basePlayer) Year() string {
	return p.year
}

func (p *basePlayer) Genre() string {
	return p.genre
}

func (p *basePlayer) Sort() int64 {
	return p.sort
}

func (p *basePlayer) Fp() (*os.File, error) {
	if p.fp != nil {
		return p.fp, nil
	}
	file, err := os.Open(p.filepath)
	if err != nil {
		return nil, errors.NewMusicNotExistError(p.name)
	}
	p.fp = file
	return p.fp, nil
}

func (p *basePlayer) CloseFp() {
	err := p.fp.Close()
	if err != nil {
		fmt.Println("关闭文件", p.name, "失败：", err)
		return
	}
	p.fp = nil
}

func (p *basePlayer) CloseStreamer() {
	err := p.streamer.Close()
	if err != nil {
		fmt.Println("关闭流媒体失败：", err)
		return
	}
	p.streamer = nil
}

func (p *basePlayer) initSpeaker() error {
	err := speaker.Init(p.format.SampleRate, p.format.SampleRate.N(time.Second/10))
	if err != nil {
		fmt.Println("初始化扬声器失败：", err)
		return err
	}
	return nil
}

func (p *basePlayer) Play() error {
	var err error
	if p.position != 0 {
		err = p.resume()
	}
	err = p.initSpeaker()
	if err != nil {
		return nil
	}
	defer speaker.Clear()
	fmt.Println("正在播放：", p.title)
	p.doPlay(p.streamer)
	return nil
}

func (p *basePlayer) doPlay(streamer beep.StreamSeekCloser) {
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (p *basePlayer) Pause() error {
	p.position = p.streamer.Position()
	speaker.Clear()
	p.CloseStreamer()
	return nil
}

func (p *basePlayer) resume() error {
	err := p.streamer.Seek(p.position)
	if err != nil {
		return err
	}
	return nil
}
