package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type IDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.IDiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorCyan, widget))
		}
		return &IDiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.IDiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
				switchType:   1,
			},
		}
	})
}

func (d *IDiamonds) getNextWidgetPos() ([][2]int, error) {
	left, right, _, _, _, midY, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midX := (right-left)/2 + left - (lib.DiamondWidth * 0.5)
	return [][2]int{
		{midX, midY - lib.DiamondHeight*2 + 1},
		{midX, midY - lib.DiamondHeight + 1},
		{midX, midY + 1},
		{midX, midY + lib.DiamondHeight + 1},
	}, nil
}

func (d *IDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	level1Y := top - lib.DiamondHeight
	return [][2]int{
		{midX, level1Y},
		{midX, level1Y - lib.DiamondHeight},
		{midX, level1Y - lib.DiamondHeight*2},
		{midX, level1Y - lib.DiamondHeight*3},
	}, nil
}

func (d *IDiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	var diamondArr = d.getDiamondCurPos()
	var switchType int
	switch d.switchType {
	case 1:
		{
			switchType = 2
			diamondArr[1][0] += lib.DiamondWidth
			diamondArr[1][1] += lib.DiamondHeight
			diamondArr[2][0] += lib.DiamondWidth * 2
			diamondArr[2][1] += lib.DiamondHeight * 2
			diamondArr[3][0] += lib.DiamondWidth * 3
			diamondArr[3][1] += lib.DiamondHeight * 3
		}
	case 2:
		{
			switchType = 1
			diamondArr[1][0] -= lib.DiamondWidth
			diamondArr[1][1] -= lib.DiamondHeight
			diamondArr[2][0] -= lib.DiamondWidth * 2
			diamondArr[2][1] -= lib.DiamondHeight * 2
			diamondArr[3][0] -= lib.DiamondWidth * 3
			diamondArr[3][1] -= lib.DiamondHeight * 3
		}
	}
	return diamondArr, switchType
}
