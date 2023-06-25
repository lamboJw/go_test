package lib

import (
	"github.com/lamboJw/gocui"
	"github.com/mattn/go-runewidth"
	"log"
	"runtime"
	"strings"
)

func SplitStringByChars(chars []rune, str string) []string {
	return strings.FieldsFunc(str, func(c rune) bool {
		for _, delimiter := range chars {
			if c == delimiter {
				return true
			}
		}
		return false
	})
}

var g *gocui.Gui

func GetGui() *gocui.Gui {
	if g == nil {
		var err error
		g, err = gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			log.Panicln(err)
		}
		if runtime.GOOS == "windows" && runewidth.IsEastAsian() {
			g.ASCII = true
		}
	}
	return g
}

func InArray(arr []interface{}, item interface{}) bool {
	for _, n := range arr {
		if n == item {
			return true
		}
	}
	return false
}

func GetViewPos(name WidgetName) (int, int, int, int, int, int, error) {
	g1 := GetGui()
	x0, y0, x1, y1, err := g1.ViewPosition(string(name))
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	left := x0 + 1
	right := x1 - 1
	top := y0 + 1
	bottom := y1 - 1
	width := (right - left) / DiamondWidth
	height := (bottom - top) / DiamondHeight
	midX := width/2*DiamondWidth + left
	midY := height/2*DiamondHeight + top
	return left, right, top, bottom, midX, midY, nil
}
