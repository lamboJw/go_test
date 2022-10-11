package lib

import (
	"fmt"
	"go_test/music_player/errors"
	"go_test/music_player/media"
	"go_test/music_player/utils"
	"math"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Manager struct {
	baseDir         string
	musicList       map[string]*media.Music
	sortedMusicList MusicList
	playingMusic    *media.Music
}

func NewManager(dir string) (*Manager, error) {
	list, err := utils.SearchFiles(dir)
	if err != nil {
		return nil, err
	}
	musicList := make(map[string]*media.Music)
	for _, path := range list {
		music, err := media.NewMusic(filepath.Join(dir, path))
		if err != nil {
			fmt.Println("初始化", path, "失败：", err)
			continue
		}
		musicList[music.Player.Id()] = music
	}
	sortedMusicList := NewMusicList(musicList)
	sort.Sort(sortedMusicList)
	return &Manager{baseDir: dir, musicList: musicList, sortedMusicList: sortedMusicList}, nil
}

func (m *Manager) GetList(name string, page int, pagesize int) MusicList {
	page = int(math.Max(float64(page), 1))
	offset := (page - 1) * pagesize
	end := offset + pagesize
	var list MusicList
	if name != "" {
		for _, item := range m.sortedMusicList {
			if strings.Contains(item.Name(), name) {
				list = append(list, item)
			}
		}
	} else {
		list = m.sortedMusicList
	}
	if end >= len(list) {
		return list[offset:]
	}
	return list[offset:end]
}

func (m *Manager) PrintList(list MusicList) {
	for k, v := range list {
		titleWidth := 20 - utils.ChineseCount(v.Title())
		artistWidth := 20 - utils.ChineseCount(v.Artist())
		albumWidth := 20 - utils.ChineseCount(v.Album())
		format := "%-5d%-" + strconv.Itoa(titleWidth) + "s%-" + strconv.Itoa(artistWidth) + "s%-" + strconv.Itoa(albumWidth) + "s\n"
		fmt.Printf(format, k+1, v.Title(), v.Artist(), v.Album())
	}
}

func (m *Manager) Play(id string, filename string) error {
	if m.musicList[id] == nil {
		return errors.NewMusicNotExistError(filename)
	}
	err := m.musicList[id].Player.Play()
	if err != nil {
		return err
	}
	return nil
}
