package lib

import "errors"

type WidgetName string

type DiamondsName string

type Direction string

const (
	HelpWidgetName  WidgetName   = "help"
	NextWidgetName  WidgetName   = "next"
	MainWidgetName  WidgetName   = "main"
	ScoreWidgetName WidgetName   = "score"
	DirectionLeft   Direction    = "left"
	DirectionDown   Direction    = "down"
	DirectionRight  Direction    = "right"
	DirectionUp     Direction    = "up"
	ODiamonds       DiamondsName = "O"
	ZDiamonds       DiamondsName = "Z"
	SDiamonds       DiamondsName = "S"
	JDiamonds       DiamondsName = "J"
	LDiamonds       DiamondsName = "L"
	IDiamonds       DiamondsName = "I"
	TDiamonds       DiamondsName = "T"
)

var DiamondsTypes = []DiamondsName{ODiamonds, ZDiamonds, SDiamonds, JDiamonds, LDiamonds, IDiamonds, TDiamonds}

const Help = `Ctrl+C：关闭程序
Enter：开始游戏
 ↑ ：旋转方块
← →：控制方块左右移动
 ↓ ：控制方块快速下落`

const (
	DiamondWidth  = 4
	DiamondHeight = 2
)

var (
	ErrNextDiamondsEmpty = errors.New("未设置下一个方块")
)
