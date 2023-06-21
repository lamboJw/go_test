package lib

import "errors"

type WidgetName string

type DiamondsName string

type Direction string

const (
	HelpWidgetName WidgetName   = "help"
	NextWidgetName WidgetName   = "next"
	MainWidgetName WidgetName   = "main"
	DirectionLeft  Direction    = "left"
	DirectionDown  Direction    = "down"
	DirectionRight Direction    = "right"
	DirectionUp    Direction    = "up"
	SquareDiamonds DiamondsName = "square"
	LeftZDiamonds  DiamondsName = "leftZ"
)

var DiamondsTypes = []DiamondsName{SquareDiamonds, LeftZDiamonds}

const Help = `Ctrl+C：Close program
Enter：开始游戏
← →：控制方块左右移动
 ↓ ：控制方块快速下落`

const (
	DiamondWidth  = 6
	DiamondHeight = 3
)

var (
	ErrNextDiamondsEmpty = errors.New("未设置下一个方块")
)
