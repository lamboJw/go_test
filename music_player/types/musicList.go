package types

import (
	"go_test/music_player/interfaces"
)

type MusicList []interfaces.MediaInfoGetter

func NewMusicList(m map[string]interfaces.MediaInfoGetterAndPlayer) MusicList {
	ms := make(MusicList, 0, len(m))
	for _, v := range m {
		ms = append(ms, v)
	}
	return ms
}

func (ms MusicList) Len() int {
	return len(ms)
}

func (ms MusicList) Less(i, j int) bool {
	return ms[i].Sort() > ms[j].Sort() // > 倒序；< 升序
}

func (ms MusicList) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
