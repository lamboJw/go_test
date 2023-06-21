package widget

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/diamonds"
	"go_test/gocui_test/lib"
	"log"
)

type NextWidget struct {
	BaseWidget
	nextDiamonds diamonds.Diamonds
}

func init() {
	creatorRegister(lib.NextWidgetName, NewNextWidget)
}

func NewNextWidget(name lib.WidgetName, x int, y int, args ...interface{}) Widget {
	width := args[0].(int)
	if width%lib.DiamondWidth != 0 {
		log.Panicln("宽度不是方块宽度的整数倍")
	}
	w := args[0].(int) + 1
	h := w / 2
	return &NextWidget{
		BaseWidget: BaseWidget{
			name: name,
			x:    x,
			y:    y,
			w:    w,
			h:    h,
		},
		nextDiamonds: nil,
	}
}

func (w *NextWidget) Layout(g *gocui.Gui) error {
	_, err := g.SetView(string(w.name), w.x, w.y, w.x+w.w+1, w.y+w.h+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	w.left, w.right, w.top, w.bottom, w.midX, w.midY, err = lib.GetViewPos(w.name)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

func (w *NextWidget) SetNextDiamonds(nextDiamonds diamonds.Diamonds) error {
	w.nextDiamonds = nextDiamonds
	return nil
}

func (w *NextWidget) GetNextDiamondsType() (lib.DiamondsName, error) {
	if w.nextDiamonds == nil {
		return "", lib.ErrNextDiamondsEmpty
	}
	return w.nextDiamonds.GetDiamondsType(), nil
}

func (w *NextWidget) DestroyNextDiamonds() error {
	if w.nextDiamonds == nil {
		return nil
	}
	if err := w.nextDiamonds.DestroyView(); err != nil {
		return err
	}
	w.nextDiamonds = nil
	return nil
}
