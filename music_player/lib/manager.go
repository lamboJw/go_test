package lib

import (
	"bufio"
	"fmt"
	"go_test/music_player/errors"
	"go_test/music_player/interfaces"
	"go_test/music_player/media"
	"go_test/music_player/types"
	"go_test/music_player/utils"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Manager struct {
	baseDir         string
	musicList       map[string]interfaces.MediaInfoGetterAndPlayer
	sortedMusicList types.MusicList
	playingMusic    interfaces.MediaInfoGetterAndPlayer
	inputChan       chan string
	runChan         chan bool
}

func NewManager(dir string) (*Manager, error) {
	list, err := utils.SearchFiles(dir)
	if err != nil {
		return nil, err
	}
	musicList := make(map[string]interfaces.MediaInfoGetterAndPlayer)
	channel := make(chan interfaces.MediaInfoGetterAndPlayer, len(list))
	for _, path := range list {
		go func(fpath string) {
			music, err := media.NewMusic(fpath)
			if err != nil {
				log.Println("初始化", fpath, "失败：", err)
				channel <- nil
			}
			channel <- music
		}(filepath.Join(dir, path))
	}
	for i := 0; i < len(list); i++ {
		music := <-channel
		if music != nil {
			musicList[music.Id()] = music
		}
	}
	sortedMusicList := types.NewMusicList(musicList)
	sort.Sort(sortedMusicList)
	return &Manager{
		baseDir:         dir,
		musicList:       musicList,
		sortedMusicList: sortedMusicList,
		inputChan:       make(chan string, 1),
		runChan:         make(chan bool, 1),
	}, nil
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
	if m.playingMusic != nil {
		if m.playingMusic == m.musicList[id] && m.playingMusic.Playing() {
			return nil
		}
		if m.playingMusic != m.musicList[id] && m.playingMusic.Playing() {
			err := m.playingMusic.Pause()
			if err != nil {
				return err
			}
		}
	}
	m.playingMusic = m.musicList[id]
	err := m.musicList[id].Play()
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Start() error {
	list := m.GetList("", 0, 20)
	m.PrintList(list)
	go func() {
		for {
			select {
			case input := <-m.inputChan:
				cmd := strings.Split(input, " ")
				switch cmd[0] {
				case "play":
					index, _ := strconv.Atoi(cmd[1])
					music := list[index-1]
					m.Play(music.Id(), music.Name())
					m.runChan <- true
				case "pause":
					if m.playingMusic != nil {
						m.playingMusic.Pause()
					}
					m.runChan <- true
				case "quit":
					if m.playingMusic != nil {
						err := m.playingMusic.Pause()
						if err != nil {
							return
						}
					}
					os.Exit(0)
				}
			}
		}
	}()
	m.runChan <- true
	for {
		<-m.runChan
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("请输入命令：")
		input, _ := reader.ReadString('\n')

		// 去除输入字符串中的换行符
		input = strings.TrimSpace(input)

		m.inputChan <- input
	}
	return nil
}
