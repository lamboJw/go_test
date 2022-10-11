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
	list := manager.GetList("", 0, 20)
	manager.PrintList(list)
	err = manager.Play(list[0].Id(), list[0].Name())
	if err != nil {
		fmt.Println(err)
	}
}
