package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type ZDiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.ZDiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorRed, widget))
		}
		return &ZDiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.ZDiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
				switchType:   1,
			},
		}
	})
}

func (d *ZDiamonds) getNextWidgetPos() ([][2]int, error) {
	_, _, top, bottom, midX, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	level2Y := (bottom-top)/2 + top
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{{midX - lib.DiamondWidth, level1Y}, {midX, level1Y}, {midX, level2Y}, {midX + lib.DiamondWidth, level2Y}}, nil
}

func (d *ZDiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{{midX - lib.DiamondWidth, level1Y}, {midX, level1Y}, {midX, level2Y}, {midX + lib.DiamondWidth, level2Y}}, nil
}

func (d *ZDiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	var diamondArr = d.getDiamondCurPos()
	var switchType int
	if d.switchType == 1 {
		switchType = 2
		diamondArr[0][1] += lib.DiamondHeight
		diamondArr[1][0] -= lib.DiamondWidth
		diamondArr[2][1] -= lib.DiamondHeight
		diamondArr[3][0] -= lib.DiamondWidth
		diamondArr[3][1] -= lib.DiamondHeight * 2
	} else {
		switchType = 1
		diamondArr[0][1] -= lib.DiamondHeight
		diamondArr[1][0] += lib.DiamondWidth
		diamondArr[2][1] += lib.DiamondHeight
		diamondArr[3][0] += lib.DiamondWidth
		diamondArr[3][1] += lib.DiamondHeight * 2
	}
	return diamondArr, switchType
}
