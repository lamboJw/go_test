package lib

import (
	"fmt"
	"go_test/music_player/errors"
	"go_test/music_player/interfaces"
	"go_test/music_player/media"
	"go_test/music_player/types"
	"go_test/music_player/utils"
	"math"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Manager struct {
	baseDir         string
	musicList       map[string]interfaces.MediaInfoGetterAndPlayer
	sortedMusicList types.MusicList
	playingMusic    *media.Music
}

func NewManager(dir string) (*Manager, error) {
	list, err := utils.SearchFiles(dir)
	if err != nil {
		return nil, err
	}
	musicList := make(map[string]interfaces.MediaInfoGetterAndPlayer)
	for _, path := range list {
		music, err := media.NewMusic(filepath.Join(dir, path))
		if err != nil {
			fmt.Println("初始化", path, "失败：", err)
			continue
		}
		musicList[music.Id()] = music
	}
	sortedMusicList := types.NewMusicList(musicList)
	sort.Sort(sortedMusicList)
	return &Manager{baseDir: dir, musicList: musicList, sortedMusicList: sortedMusicList}, nil
}

func (m *Manager) GetList(name string, page int, pagesize int) types.MusicList {
	page = int(math.Max(float64(page), 1))
	offset := (page - 1) * pagesize
	end := offset + pagesize
	var list types.MusicList
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

func (m *Manager) PrintList(list types.MusicList) {
	fmt.Printf("%-4s%-18s%-18s%-47s%-32s\n", "序号", "标题", "歌手", "专辑", "ID")
	for k, v := range list {
		titleWidth := 20 - utils.ChineseCount(v.Title())
		artistWidth := 20 - utils.ChineseCount(v.Artist())
		albumWidth := 50 - utils.ChineseCount(v.Album())
		format := "%-5d%-" + strconv.Itoa(titleWidth) + "s%-" + strconv.Itoa(artistWidth) + "s%-" + strconv.Itoa(albumWidth) + "s%-32s\n"
		fmt.Printf(format, k+1, v.Title(), v.Artist(), v.Album(), v.Id())
	}
}

func (m *Manager) Play(id string, filename string) error {
	if m.musicList[id] == nil {
		return errors.NewMusicNotExistError(filename)
	}
	err := m.musicList[id].Play()
	if err != nil {
		return err
	}
	return nil
}
