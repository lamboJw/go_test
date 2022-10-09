package main

import (
	"fmt"
	"go_test/music_player/lib"
)

func main() {
	manager, err := lib.NewManager("E:\\CloudMusic")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range manager.GetList(0, 3) {
		fmt.Printf("path: %d, info：%+v\n", k, v.Player)
	}
}
