package lib

import (
	"fmt"
	"go_test/music_player/errors"
	"go_test/music_player/media"
	"go_test/music_player/utils"
	"math"
	"path/filepath"
	"sort"
	"strings"
)

type Manager struct {
	baseDir         string
	musicList       map[string]*media.Music
	sortedMusicList MusicListSorter
	playingMusic    media.Music
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
	sortedMusicList := NewMusicListSorter(musicList)
	sort.Sort(sortedMusicList)
	return &Manager{baseDir: dir, musicList: musicList, sortedMusicList: sortedMusicList}, nil
}

func (m *Manager) GetList(page int, pagesize int) []*media.Music {
	page = int(math.Max(float64(page), 1))
	offset := (page - 1) * pagesize
	end := offset + pagesize
	fmt.Println(offset, end, len(m.sortedMusicList))
	if end >= len(m.sortedMusicList) {
		return m.sortedMusicList[offset:]
	}
	return m.sortedMusicList[offset:end]
}

func (m *Manager) Search(name string, page int, pagesize int) []*media.Music {
	offset := (page - 1) * pagesize
	end := offset + pagesize
	if end >= len(m.sortedMusicList) {
		end = -1
	}
	var list []*media.Music
	for _, item := range m.sortedMusicList {
		if strings.Contains(item.Player.Name(), name) {
			list = append(list, item)
		}
	}
	return list[offset:end]
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
