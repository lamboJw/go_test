package widget

import "go_test/gocui_test/lib"

type Widget interface {
	GetSetter
}

type GetSetter interface {
	Getter
	Setter
}

type Getter interface {
	Name() lib.WidgetName
	X() int
	Y() int
	W() int
	H() int
	Left() (int, error)
	Right() (int, error)
	Top() (int, error)
	Bottom() (int, error)
	MidX() (int, error)
	MidY() (int, error)
}

type Setter interface {
	SetLeft(left int)
	SetRight(right int)
	SetTop(top int)
	SetBottom(bottom int)
	SetMidX(midX int)
	SetMidY(midY int)
}
