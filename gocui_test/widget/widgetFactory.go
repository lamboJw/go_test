package widget

import (
	"errors"
	"go_test/gocui_test/lib"
	"log"
)

type Creator func(name lib.WidgetName, x int, y int, args ...interface{}) Widget

var typeMap = make(map[lib.WidgetName]Creator)

var next *NextWidget
var help *HelpWidget
var main *MainWidget
var score *ScoreWidget

/*
注册工厂方法
*/
func creatorRegister(typeStr lib.WidgetName, factory Creator) {
	typeMap[typeStr] = factory
}

// Create 根据类型调用对应的工厂方法
func Create(typeStr lib.WidgetName, x int, y int, args ...interface{}) (Widget, error) {
	if factory, ok := typeMap[typeStr]; ok {
		return factory(typeStr, x, y, args...), nil
	} else {
		return nil, errors.New(string("没有找到" + typeStr + "类型的组件"))
	}
}

var nextWidgetWidth = 5
var nextWidgetHeight = 5
var mainWidgetWidth = 10
var mainWidgetHeight = 15

func GetMainWidget() *MainWidget {
	if main == nil {
		_widget, err := Create(lib.MainWidgetName, lib.DiamondWidth*nextWidgetWidth+5, 0, lib.DiamondWidth*mainWidgetWidth, lib.DiamondHeight*mainWidgetHeight)
		if err != nil {
			log.Panicln(err)
		}
		main = _widget.(*MainWidget)
	}
	return main
}

func GetNextWidget() *NextWidget {
	if next == nil {
		_widget, err := Create(lib.NextWidgetName, 0, 0, lib.DiamondWidth*nextWidgetWidth)
		if err != nil {
			log.Panicln(err)
		}
		next = _widget.(*NextWidget)
	}
	return next
}

func GetHelpWidget() *HelpWidget {
	if help == nil {
		_widget, err := Create(lib.HelpWidgetName, lib.DiamondWidth*(mainWidgetWidth+nextWidgetWidth)+10, 0, lib.Help)
		if err != nil {
			log.Panicln(err)
		}
		help = _widget.(*HelpWidget)
	}
	return help
}

func GetScoreWidget() *ScoreWidget {
	if score == nil {
		_widget, err := Create(lib.ScoreWidgetName, 0, lib.DiamondHeight*nextWidgetHeight+3, lib.DiamondWidth*nextWidgetWidth)
		if err != nil {
			log.Panicln(err)
		}
		score = _widget.(*ScoreWidget)
	}
	return score
}
