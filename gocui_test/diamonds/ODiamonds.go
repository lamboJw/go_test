package diamonds

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"strconv"
)

type ODiamonds struct {
	BaseDiamonds
}

/*
引入包的时候，会自动调用init方法
*/
func init() {
	creatorRegister(lib.ODiamonds, func(index int, widget lib.WidgetName) Diamonds {
		var arr []*Diamond
		indexStr := strconv.Itoa(index)
		for _, i := range []string{"0", "1", "2", "3"} {
			arr = append(arr, NewDiamond(indexStr+"_"+i, gocui.ColorYellow, widget))
		}
		return &ODiamonds{
			BaseDiamonds: BaseDiamonds{
				diamondsType: lib.ODiamonds,
				diamondArr:   arr,
				index:        index,
				widget:       widget,
			},
		}
	})
}

func (d *ODiamonds) getNextWidgetPos() ([][2]int, error) {
	left, right, top, bottom, _, _, err := lib.GetViewPos(lib.NextWidgetName)
	if err != nil {
		return nil, err
	}
	midX := (right-left)/2 + left
	midY := (bottom-top)/2 + top
	x := midX - lib.DiamondWidth
	y := midY - lib.DiamondHeight
	return [][2]int{{x, y}, {midX, y}, {x, midY}, {midX, midY}}, nil
}

func (d *ODiamonds) getMainWidgetPos() ([][2]int, error) {
	_, _, top, _, midX, _, err := lib.GetViewPos(lib.MainWidgetName)
	if err != nil {
		return nil, err
	}
	x := midX - lib.DiamondWidth
	level2Y := top - lib.DiamondHeight
	level1Y := level2Y - lib.DiamondHeight
	return [][2]int{{x, level1Y}, {midX, level1Y}, {x, level2Y}, {midX, level2Y}}, nil
}

func (d *ODiamonds) GetSwitchDirectionPos() ([][2]int, int) {
	return d.getDiamondCurPos(), d.switchType
}
