package lib

import "go_test/music_player/media"

type MusicListSorter []*media.Music

func NewMusicListSorter(m map[string]*media.Music) MusicListSorter {
	ms := make(MusicListSorter, 0, len(m))
	for _, v := range m {
		ms = append(ms, v)
	}
	return ms
}

func (ms MusicListSorter) Len() int {
	return len(ms)
}

func (ms MusicListSorter) Less(i, j int) bool {
	return ms[i].Player.Sort() > ms[j].Player.Sort() // > 倒序；< 升序
}

func (ms MusicListSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
