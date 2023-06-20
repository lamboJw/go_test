package diamonds

import "go_test/gocui_test/lib"

type Diamonds interface {
	Initializer
	Eventer
}

type Initializer interface {
	GetWidgetName() lib.WidgetName
	getNextWidgetPos() ([][2]int, error)
	getMainWidgetPos() ([][2]int, error)
	DrawDiamonds(pos [][2]int) error
}

type Eventer interface {
	GetDiamondArr() []*Diamond
	MoveDown() error
	MoveLeft() error
	MoveRight() error
	SwitchDirection() error
	DestroyView()
}
