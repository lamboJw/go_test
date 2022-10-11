package media

import (
	"encoding/binary"
	"go_test/music_player/errors"
	"go_test/music_player/types"
	"go_test/music_player/utils"
	"golang.org/x/text/encoding/unicode"
	"os"
	"strings"
)

type ID3 interface {
	getTag(fp *os.File) error
	Title() string
	Artist() string
	Album() string
	Year() string
	Genre() string
}

func newID3(fp *os.File) (ID3, error) {
	var id3 ID3
	id3 = &ID3V2{}
	err := id3.getTag(fp)
	if err != nil {
		id3 = &ID3V1{}
		err = id3.getTag(fp)
		if err != nil {
			return nil, err
		} else {
			return id3, nil
		}
	} else {
		return id3, nil
	}
}

type ID3V1 struct {
	title  string
	artist string
	album  string
	year   string
	genre  types.Genre
}

func (i *ID3V1) Title() string {
	return i.title
}

func (i *ID3V1) Artist() string {
	return i.artist
}

func (i *ID3V1) Album() string {
	return i.album
}

func (i *ID3V1) Year() string {
	return i.year
}

func (i *ID3V1) Genre() string {
	return i.genre.String()
}

func (i *ID3V1) getTag(fp *os.File) error {
	fInfo, _ := fp.Stat()
	fileSize := fInfo.Size()
	id3 := make([]byte, 128)
	_, err := fp.ReadAt(id3, fileSize-128)
	if err != nil {
		return err
	}
	if string(id3[:3]) != "TAG" {
		return errors.NewMusicTagError(fp.Name())
	}
	i.title = strings.Trim(string(id3[3:33]), "\x00")
	i.album = strings.Trim(string(id3[33:63]), "\x00")
	i.artist = strings.Trim(string(id3[63:93]), "\x00")
	i.year = string(id3[93:97])
	i.genre = types.Genre(id3[127])
	return nil
}

type ID3V2 struct {
	curOffset int64
	title     string
	artist    string
	album     string
	year      string
	genre     string
}

func (i *ID3V2) Genre() string {
	return i.genre
}

func (i *ID3V2) Year() string {
	return i.year
}

func (i *ID3V2) Album() string {
	return i.album
}

func (i *ID3V2) Artist() string {
	return i.artist
}

func (i *ID3V2) Title() string {
	return i.title
}

func (i *ID3V2) getTag(fp *os.File) error {
	header := &SignHeader{}
	err := header.init(fp)
	if err != nil {
		return err
	}
	i.curOffset = 10
	for i.curOffset < header.Size {
		content, err := i.getFrameContent(fp)
		if err != nil {
			if _, ok := err.(*errors.MusicTagEmptyError); ok {
				break
			}
			return err
		}
		switch content["frameId"] {
		case "TIT2":
			i.title = content["content"]
		case "TPE1":
			i.artist = content["content"]
		case "TALB":
			i.album = content["content"]
		case "TYER":
			i.year = content["content"]
		case "TCON":
			i.genre = content["content"]
		}
	}
	return nil
}

type SignHeader struct {
	Version  uint8
	Revision uint8
	Flag     map[string]uint8
	Size     int64
}

func (h *SignHeader) init(fp *os.File) error {
	header, err := h.getSignHeader(fp)
	if err != nil {
		return err
	}
	h.Version = header[3:4][0]
	h.Revision = header[4:5][0]
	flag := header[5:6][0]
	h.Flag = make(map[string]uint8, 3)
	h.Flag["a"] = utils.GetOffsetBit(flag, 8)
	h.Flag["b"] = utils.GetOffsetBit(flag, 7)
	h.Flag["c"] = utils.GetOffsetBit(flag, 6)
	size := header[6:]
	h.Size = int64(size[0])*0x200000 + int64(size[1])*0x4000 + int64(size[2])*0x80 + int64(size[3])
	return nil
}

func (h *SignHeader) getSignHeader(fp *os.File) ([]byte, error) {
	signHeader := make([]byte, 10)
	_, err := fp.Read(signHeader)
	if err != nil {
		return nil, err
	}
	if string(signHeader[:3]) != "ID3" {
		return nil, errors.NewMusicTagError(fp.Name())
	}
	return signHeader, nil
}

func (h *SignHeader) getSignHeaderSize(signHeader []byte) uint32 {
	size := signHeader[6:]
	//fmt.Println(biu.ToBinaryString(size))
	totalSize := make([]byte, 4)
	totalSize[3] = utils.SetOffsetBit(size[3], 8, utils.GetOffsetBit(size[2], 1))
	size2 := utils.SetOffsetBit(size[2]>>1, 7, utils.GetOffsetBit(size[1], 1))
	totalSize[2] = utils.SetOffsetBit(size2, 8, utils.GetOffsetBit(size[1], 2))
	size1 := utils.SetOffsetBit(size[1]>>2, 6, utils.GetOffsetBit(size[0], 1))
	size1 = utils.SetOffsetBit(size1, 7, utils.GetOffsetBit(size[0], 2))
	totalSize[1] = utils.SetOffsetBit(size1, 8, utils.GetOffsetBit(size[0], 3))
	totalSize[0] = size[0] >> 3
	//fmt.Println(biu.ToBinaryString(totalSize))
	return binary.BigEndian.Uint32(totalSize)
}

func (i *ID3V2) getFrameContent(fp *os.File) (map[string]string, error) {
	frameHeader := make([]byte, 10)
	_, err := fp.ReadAt(frameHeader, i.curOffset)
	if err != nil {
		return nil, err
	}
	i.curOffset += 10
	content := make(map[string]string, 3)
	frameId := string(frameHeader[:4])
	content["frameId"] = frameId
	size := frameHeader[4:8]
	totalSize := int64(size[0])*0x100000 + int64(size[1])*0x10000 + int64(size[2])*0x100 + int64(size[3])
	if totalSize == 0 {
		return nil, errors.NewMusicTagEmptyError(fp.Name(), i.curOffset)
	}
	encoding := make([]byte, 1)
	_, err = fp.ReadAt(encoding, i.curOffset)
	if err != nil {
		return nil, err
	}
	frameContent := make([]byte, totalSize-1)
	_, err = fp.ReadAt(frameContent, i.curOffset+1)
	if err != nil {
		return nil, err
	}
	if encoding[0] == 1 || encoding[0] == 2 {
		var endianness unicode.Endianness
		if encoding[0] == 1 {
			endianness = unicode.LittleEndian
		} else {
			endianness = unicode.BigEndian
		}
		decoder := unicode.UTF16(endianness, unicode.IgnoreBOM).NewDecoder()
		frameContent, err = decoder.Bytes(frameContent)
		if err != nil {
			return nil, err
		}
	}
	content["content"] = string(frameContent)
	i.curOffset += totalSize
	return content, nil
}
