package media

import (
	"github.com/mewkiz/flac/meta"
	"go_test/music_player/errors"
	"go_test/music_player/types"
	"os"
)

func newFlacMetaData(fp *os.File) (*FlacMetaData, error) {
	var meta = &FlacMetaData{}
	err := meta.getTag(fp)
	return meta, err
}

type FlacMetaData struct {
	title    string
	artist   string
	album    string
	year     string
	genre    types.Genre
	metaData map[meta.Type]*meta.Block
}

func (f *FlacMetaData) Title() string {
	return f.title
}

func (f *FlacMetaData) Artist() string {
	return f.artist
}

func (f *FlacMetaData) Album() string {
	return f.album
}

func (f *FlacMetaData) Year() string {
	return f.year
}

func (f *FlacMetaData) Genre() string {
	return f.genre.String()
}

func (f *FlacMetaData) getTag(fp *os.File) error {
	flacFlag := make([]byte, 4)
	_, err := fp.Read(flacFlag)
	if err != nil {
		return err
	}
	if string(flacFlag) != "fLaC" {
		return errors.NewMusicTagError(fp.Name())
	}
	f.metaData = make(map[meta.Type]*meta.Block)
	for {
		block, err := meta.Parse(fp)
		if err != nil {
			return err
		}
		if block.Type == meta.TypeVorbisComment {
			tags := block.Body.(*meta.VorbisComment).Tags
			for _, tag := range tags {
				switch tag[0] {
				case "title":
					f.title = tag[1]
				case "artist":
					f.artist = tag[1]
				case "album":
					f.album = tag[1]
				case "date":
					f.year = tag[1]
				}
			}
			break
		}
		if block.IsLast {
			break
		}
	}
	return nil
}
