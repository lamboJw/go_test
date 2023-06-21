package diamonds

import (
	"fmt"
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"log"
)

type Diamond struct {
	name   string
	x      int
	y      int
	color  gocui.Attribute
	widget lib.WidgetName
}

func (d *Diamond) String() string {
	return fmt.Sprintf("%s: x:%d y:%d color:%v widget:%s\n", d.name, d.x, d.y, d.color, d.widget)
}

func NewDiamond(name string, color gocui.Attribute, widget lib.WidgetName) *Diamond {
	return &Diamond{
		name:   name,
		color:  color,
		widget: widget,
	}
}

func (d *Diamond) Destroy() error {
	g := lib.GetGui()
	if err := g.DeleteView(d.name); err != nil && err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func (d *Diamond) setView() error {
	g := lib.GetGui()
	if d.widget == lib.MainWidgetName {
		_, _, top, _, _, _, err := lib.GetViewPos(lib.MainWidgetName)
		if err != nil && err != gocui.ErrUnknownView {
			log.Println(d.name + "get main view position error")
			return err
		}
		if d.y < top {
			if err = d.Destroy(); err != nil {
				log.Println(d.name, "destroy less than top")
				return err
			}
			return nil
		}
	}
	view, err := g.SetView(d.name, d.x, d.y, d.x+lib.DiamondWidth, d.y+lib.DiamondHeight)
	if err != nil && err != gocui.ErrUnknownView {
		log.Println(d.name + "set view error")
		return err
	}
	view.Frame = true
	view.BgColor = d.color
	view.FgColor = d.color
	return nil
}

func (d *Diamond) Init(x int, y int) error {
	d.x = x
	d.y = y
	if err := d.setView(); err != nil {
		return err
	}
	return nil
}

func (d *Diamond) MoveLeft() error {
	d.x = d.x - lib.DiamondWidth
	if err := d.setView(); err != nil {
		return err
	}
	return nil
}

func (d *Diamond) MoveRight() error {
	d.x = d.x + lib.DiamondWidth
	if err := d.setView(); err != nil {
		return err
	}
	return nil
}

func (d *Diamond) MoveDown() error {
	d.y = d.y + lib.DiamondHeight
	if err := d.setView(); err != nil {
		return err
	}
	return nil
}

func (d *Diamond) GetPos() [2]int {
	return [2]int{d.x, d.y}
}
