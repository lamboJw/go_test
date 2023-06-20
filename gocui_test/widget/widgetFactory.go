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

func GetMainWidget() *MainWidget {
	if main == nil {
		_widget, err := Create(lib.MainWidgetName, 35, 0, lib.DiamondWidth*20, lib.DiamondHeight*17)
		if err != nil {
			log.Panicln(err)
		}
		main = _widget.(*MainWidget)
	}
	return main
}

func GetNextWidget() *NextWidget {
	if next == nil {
		_widget, err := Create(lib.NextWidgetName, 0, 0, lib.DiamondWidth*5)
		if err != nil {
			log.Panicln(err)
		}
		next = _widget.(*NextWidget)
	}
	return next
}

func GetHelpWidget() *HelpWidget {
	if help == nil {
		g1 := lib.GetGui()
		maxX, _ := g1.Size()
		_widget, err := Create(lib.HelpWidgetName, maxX-40, 0, lib.Help)
		if err != nil {
			log.Panicln(err)
		}
		help = _widget.(*HelpWidget)
	}
	return help
}
