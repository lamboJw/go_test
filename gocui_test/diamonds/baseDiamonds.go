package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
)

type BaseDiamonds struct {
	diamondArr []*Diamond
	x, y       int
	index      int
	widget     lib.WidgetName
	switchType int
}

func (d *BaseDiamonds) DrawDiamonds(pos [][2]int) error {
	for index, item := range d.diamondArr {
		if err := item.Init(pos[index][0], pos[index][1]); err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) MoveDown() error {
	for _, diamond := range d.diamondArr {
		err := diamond.MoveDown()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) MoveLeft() error {
	for _, diamond := range d.diamondArr {
		err := diamond.MoveLeft()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) MoveRight() error {
	for _, diamond := range d.diamondArr {
		err := diamond.MoveRight()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) RefreshDiamonds() error {
	for _, diamond := range d.diamondArr {
		err := diamond.setView()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) GetDiamondArr() []*Diamond {
	return d.diamondArr
}

func (d *BaseDiamonds) GetWidgetName() lib.WidgetName {
	return d.widget
}

func (d *BaseDiamonds) DestroyView() {
	g := lib.GetGui()
	g.Update(func(gui *gocui.Gui) error {
		for _, diamond := range d.diamondArr {
			if err := diamond.Destroy(); err != nil {
				return err
			}
		}
		return nil
	})
}
