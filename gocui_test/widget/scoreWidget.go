package widget

import (
	"fmt"
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"log"
)

type ScoreWidget struct {
	BaseWidget
	score int
	level int
}

func init() {
	creatorRegister(lib.ScoreWidgetName, func(name lib.WidgetName, x, y int, args ...interface{}) Widget {
		w := args[0].(int)
		return &ScoreWidget{
			BaseWidget: BaseWidget{
				name: name,
				x:    x,
				y:    y,
				w:    w,
				h:    2,
			},
		}
	})
}

func (w *ScoreWidget) Layout(g *gocui.Gui) error {
	_, err := g.SetView(string(w.name), w.x, w.y, w.x+w.w+2, w.y+w.h+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	w.updateScore()
	return nil
}

func (w *ScoreWidget) updateScore() {
	v, err := lib.GetGui().View(string(lib.ScoreWidgetName))
	if err != nil {
		log.Panicln(err)
	}
	v.Clear()
	fmt.Fprintf(v, "分数：%d\n等级：%d", main.score, main.level)
}
