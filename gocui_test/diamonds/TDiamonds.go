package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type TDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.TDiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorMagenta, widget))
		}
		return &TDiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.TDiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
				switchType:   1,
			},
		}
	})
}

func (d *TDiamonds) getNextWidgetPos() ([][2]int, error) {
	left, right, top, bottom, _, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midX := (right-left)/2 + left
	level2Y := (bottom-top)/2 + top
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth*1.5, level2Y},
		{midX - lib.DiamondWidth*0.5, level2Y},
		{midX + lib.DiamondWidth*0.5, level2Y},
		{midX - lib.DiamondWidth*0.5, level1Y},
	}, nil
}

func (d *TDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth, level2Y},
		{midX, level2Y},
		{midX + lib.DiamondWidth, level2Y},
		{midX, level1Y},
	}, nil
}

func (d *TDiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	var diamondArr = d.getDiamondCurPos()
	var switchType int
	switch d.switchType {
	case 1:
		{
			switchType = 2
			diamondArr[0][0] += lib.DiamondWidth
			diamondArr[0][1] -= lib.DiamondHeight
			diamondArr[2][0] -= lib.DiamondWidth
			diamondArr[2][1] += lib.DiamondHeight
			diamondArr[3][0] += lib.DiamondWidth
			diamondArr[3][1] += lib.DiamondHeight
		}
	case 2:
		{
			switchType = 3
			diamondArr[0][0] += lib.DiamondWidth
			diamondArr[0][1] += lib.DiamondHeight
			diamondArr[2][0] -= lib.DiamondWidth
			diamondArr[2][1] -= lib.DiamondHeight
			diamondArr[3][0] -= lib.DiamondWidth
			diamondArr[3][1] += lib.DiamondHeight
		}
	case 3:
		{
			switchType = 4
			diamondArr[0][0] -= lib.DiamondWidth
			diamondArr[0][1] += lib.DiamondHeight
			diamondArr[2][0] += lib.DiamondWidth
			diamondArr[2][1] -= lib.DiamondHeight
			diamondArr[3][0] -= lib.DiamondWidth
			diamondArr[3][1] -= lib.DiamondHeight
		}
	case 4:
		{
			switchType = 1
			diamondArr[0][0] -= lib.DiamondWidth
			diamondArr[0][1] -= lib.DiamondHeight
			diamondArr[2][0] += lib.DiamondWidth
			diamondArr[2][1] += lib.DiamondHeight
			diamondArr[3][0] += lib.DiamondWidth
			diamondArr[3][1] -= lib.DiamondHeight
		}
	}
	return diamondArr, switchType
}
