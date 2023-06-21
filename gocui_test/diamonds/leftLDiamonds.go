package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type LeftLDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.LeftLDiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3", "4"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorGreen, widget))
		}
		return &LeftLDiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.LeftLDiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
				switchType:   1,
			},
		}
	})
}

func (d *LeftLDiamonds) getNextWidgetPos() ([][2]int, error) {
	left, right, top, bottom, _, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midX := (right-left)/2 + left
	level2Y := (bottom-top)/2 + top
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth*2, level1Y},
		{midX - lib.DiamondWidth*2, level2Y},
		{midX - lib.DiamondWidth*1, level2Y},
		{midX, level2Y},
		{midX + lib.DiamondWidth, level2Y},
	}, nil
}

func (d *LeftLDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth*2, level1Y},
		{midX - lib.DiamondWidth*2, level2Y},
		{midX - lib.DiamondWidth*1, level2Y},
		{midX, level2Y},
		{midX + lib.DiamondWidth, level2Y},
	}, nil
}

func (d *LeftLDiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	var diamondArr = d.getDiamondCurPos()
	var switchType int
	switch d.switchType {
	case 1:
		{
			switchType = 2
			diamondArr[0][0] += lib.DiamondWidth * 2
			diamondArr[0][1] -= lib.DiamondHeight
			diamondArr[1][0] += lib.DiamondWidth
			diamondArr[1][1] -= lib.DiamondHeight * 2
			diamondArr[2][1] -= lib.DiamondHeight
			diamondArr[3][0] -= lib.DiamondWidth
			diamondArr[4][0] -= lib.DiamondWidth * 2
			diamondArr[4][1] += lib.DiamondHeight
		}
	case 2:
		{
			switchType = 3
		}
	case 3:
		{
			switchType = 4
		}
	case 4:
		{
			switchType = 1
		}
	}
	return diamondArr, switchType
}
