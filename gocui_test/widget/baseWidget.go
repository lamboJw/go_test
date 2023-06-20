package widget

import (
	"go_test/gocui_test/lib"
)

type BaseWidget struct {
	name                     lib.WidgetName
	x, y, w, h               int
	left, right, top, bottom int
	midX, midY               int
}

func (b *BaseWidget) SetLeft(left int) {
	b.left = left
}

func (b *BaseWidget) SetRight(right int) {
	b.right = right
}

func (b *BaseWidget) SetTop(top int) {
	b.top = top
}

func (b *BaseWidget) SetBottom(bottom int) {
	b.bottom = bottom
}

func (b *BaseWidget) SetMidX(midX int) {
	b.midX = midX
}

func (b *BaseWidget) SetMidY(midY int) {
	b.midY = midY
}

func (b BaseWidget) X() int {
	return b.x
}

func (b BaseWidget) Y() int {
	return b.y
}

func (b BaseWidget) W() int {
	return b.w
}

func (b BaseWidget) H() int {
	return b.h
}

func (b BaseWidget) Left() (int, error) {
	return b.left, nil
}

func (b BaseWidget) Right() (int, error) {
	return b.right, nil
}

func (b BaseWidget) Top() (int, error) {
	return b.top, nil
}

func (b BaseWidget) Bottom() (int, error) {
	return b.bottom, nil
}

func (b BaseWidget) MidX() (int, error) {
	return b.midX, nil
}

func (b BaseWidget) MidY() (int, error) {
	return b.midY, nil
}

func (b BaseWidget) Name() lib.WidgetName {
	return b.name
}
