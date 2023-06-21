package diamonds

import "go_test/gocui_test/lib"

type Diamonds interface {
	Eventer
	Getter
}

type Getter interface {
	GetWidgetName() lib.WidgetName
	getNextWidgetPos() ([][2]int, error)
	getMainWidgetPos() ([][2]int, error)
	GetDiamondsType() lib.DiamondsName
	GetDiamondArr() []*Diamond
	GetSwitchDirectionPos() ([][2]int, int)
}

type Eventer interface {
	MoveDown() error
	MoveLeft() error
	MoveRight() error
	SwitchDirection(diamondArr [][2]int, switchType int) error
	DestroyView() error
	DrawDiamonds(pos [][2]int) error
}
