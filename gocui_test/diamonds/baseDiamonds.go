package diamonds

import (
	"go_test/gocui_test/lib"
)

type BaseDiamonds struct {
	diamondsType lib.DiamondsName
	diamondArr   []*Diamond
	x, y         int
	index        int
	widget       lib.WidgetName
	switchType   int
}

func (d BaseDiamonds) GetDiamondsType() lib.DiamondsName {
	return d.diamondsType
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

func (d *BaseDiamonds) DestroyView() error {
	for _, diamond := range d.diamondArr {
		if err := diamond.Destroy(); err != nil {
			return err
		}
	}
	return nil
}

func (d *BaseDiamonds) getDiamondCurPos() [][2]int {
	var diamondArr = make([][2]int, len(d.diamondArr))
	for k, diamond := range d.diamondArr {
		diamondArr[k] = diamond.GetPos()
	}
	return diamondArr
}

func (d *BaseDiamonds) SwitchDirection(diamondArr [][2]int, switchType int) error {
	d.switchType = switchType
	for k, diamond := range diamondArr {
		d.diamondArr[k].x = diamond[0]
		d.diamondArr[k].y = diamond[1]
	}
	if err := d.RefreshDiamonds(); err != nil {
		return err
	}
	return nil
}
