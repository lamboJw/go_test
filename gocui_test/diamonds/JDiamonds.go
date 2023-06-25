package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type JDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.JDiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorBlue, widget))
		}
		return &JDiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.JDiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
				switchType:   1,
			},
		}
	})
}

func (d *JDiamonds) getNextWidgetPos() ([][2]int, error) {
	left, right, top, bottom, _, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midX := (right-left)/2 + left
	level2Y := (bottom-top)/2 + top
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth*1.5, level1Y},
		{midX - lib.DiamondWidth*1.5, level2Y},
		{midX - lib.DiamondWidth*0.5, level2Y},
		{midX + lib.DiamondWidth*0.5, level2Y},
	}, nil
}

func (d *JDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{
		{midX - lib.DiamondWidth, level1Y},
		{midX - lib.DiamondWidth, level2Y},
		{midX, level2Y},
		{midX + lib.DiamondWidth, level2Y},
	}, nil
}

func (d *JDiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	var diamondArr = d.getDiamondCurPos()
	var switchType int
	switch d.switchType {
	case 1:
		{
			switchType = 2
			diamondArr[0][0] += lib.DiamondWidth
			diamondArr[1][1] -= lib.DiamondHeight
			diamondArr[2][0] -= lib.DiamondWidth
			diamondArr[3][0] -= lib.DiamondWidth * 2
			diamondArr[3][1] += lib.DiamondHeight
		}
	case 2:
		{
			switchType = 3
			diamondArr[0][1] += lib.DiamondHeight
			diamondArr[1][0] += lib.DiamondWidth
			diamondArr[2][1] -= lib.DiamondHeight
			diamondArr[3][0] -= lib.DiamondWidth
			diamondArr[3][1] -= lib.DiamondHeight * 2
		}
	case 3:
		{
			switchType = 4
			diamondArr[0][0] -= lib.DiamondWidth
			diamondArr[1][1] += lib.DiamondHeight
			diamondArr[2][0] += lib.DiamondWidth
			diamondArr[3][0] += lib.DiamondWidth * 2
			diamondArr[3][1] -= lib.DiamondHeight
		}
	case 4:
		{
			switchType = 1
			diamondArr[0][1] -= lib.DiamondHeight
			diamondArr[1][0] -= lib.DiamondWidth
			diamondArr[2][1] += lib.DiamondHeight
			diamondArr[3][0] += lib.DiamondWidth
			diamondArr[3][1] += lib.DiamondHeight * 2
		}
	}
	return diamondArr, switchType
}
