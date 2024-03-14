package music_player

import (
	"fmt"
	"go_test/music_player/lib"
)

func TestPlayer() {
	manager, err := lib.NewManager("D:\\CloudMusic")
	if err != nil {
		fmt.Println(err)
	}
	manager.Start()
}
