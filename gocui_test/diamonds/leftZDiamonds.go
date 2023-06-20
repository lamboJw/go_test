package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type LeftZDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.LeftZDiamonds, NewLeftZDiamonds)
}

func NewLeftZDiamonds(index int, widget lib.WidgetName) Diamonds {
	var arr []*Diamond
	indexStr := strconv.Itoa(index)
	for _, i := range []string{"0", "1", "2", "3"} {
		arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorMagenta, widget))
	}
	return &LeftZDiamonds{
		BaseDiamonds: BaseDiamonds{
			diamondArr: arr,
			index:      index,
			widget:     widget,
			switchType: 1,
		},
	}
}

func (d *LeftZDiamonds) getNextWidgetPos() ([][2]int, error) {
	_, _, top, bottom, midX, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midY := (bottom-top)/2 + top
	x := midX - lib.DiamondWidth
	y := midY - lib.DiamondHeight
	return [][2]int{{x, y}, {midX, y}, {midX, midY}, {midX + lib.DiamondWidth, midY}}, nil
}

func (d *LeftZDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	x := midX - lib.DiamondWidth
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{{x, level1Y}, {midX, level1Y}, {midX, level2Y}, {midX + lib.DiamondWidth, level2Y}}, nil
}

func (d *LeftZDiamonds) SwitchDirection() error {
	if d.switchType == 1 {
		d.switchType = 2
		d.diamondArr[0].y += lib.DiamondHeight
		d.diamondArr[1].x -= lib.DiamondWidth
		d.diamondArr[2].y -= lib.DiamondHeight
		d.diamondArr[3].x -= lib.DiamondWidth
		d.diamondArr[3].y -= lib.DiamondHeight * 2
	} else {
		d.switchType = 1
		d.diamondArr[0].y -= lib.DiamondHeight
		d.diamondArr[1].x += lib.DiamondWidth
		d.diamondArr[2].y += lib.DiamondHeight
		d.diamondArr[3].x += lib.DiamondWidth
		d.diamondArr[3].y += lib.DiamondHeight * 2
	}
	if err := d.RefreshDiamonds(); err != nil {
		return err
	}
	return nil
}
