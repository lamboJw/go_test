package diamonds

import (
	"errors"
	"go_test/gocui_test/lib"
	"log"
	"math/rand"
	"time"
)

type Creator func(index int, widget lib.WidgetName) Diamonds

var typeMap = make(map[lib.DiamondsName]Creator)

/*
注册工厂方法
*/
func creatorRegister(typeStr lib.DiamondsName, factory Creator) {
	typeMap[typeStr] = factory
}

func initDiamonds(d Initializer) error {
	if d.GetWidgetName() != lib.NextWidgetName && d.GetWidgetName() != lib.MainWidgetName {
		log.Panicln("组件名称错误")
	}
	var pos [][2]int
	var err error
	if d.GetWidgetName() == lib.NextWidgetName {
		pos, err = d.getNextWidgetPos()
		if err != nil {
			return err
		}
	} else {
		pos, err = d.getMainWidgetPos()
		if err != nil {
			return err
		}
	}
	if err = d.DrawDiamonds(pos); err != nil {
		return err
	}
	return nil
}

// Create 根据类型调用对应的工厂方法
func Create(typeStr lib.DiamondsName, index int, widget lib.WidgetName) (Eventer, error) {
	if factory, ok := typeMap[typeStr]; ok {
		diamond := factory(index, widget)
		if err := initDiamonds(diamond); err != nil {
			return nil, err
		}
		return diamond, nil
	} else {
		return nil, errors.New(string("没有找到" + typeStr + "类型的方块"))
	}
}

func GetRandomType() lib.DiamondsName {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(lib.DiamondsTypes))
	return lib.DiamondsTypes[randIndex]
}
