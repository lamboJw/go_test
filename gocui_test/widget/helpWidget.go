package widget

import (
	"fmt"
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/diamonds"
	"go_test/gocui_test/lib"
	"strconv"
	"strings"
)

type HelpWidget struct {
	BaseWidget
	body         string
	keyWidth     int
	descWidth    int
	splitBody    []map[string]string
	nextDiamonds *diamonds.Diamonds
}

func init() {
	creatorRegister(lib.HelpWidgetName, func(name lib.WidgetName, x, y int, args ...interface{}) Widget {
		body := args[0].(string)
		lines := strings.Split(body, "\n")

		w := 0
		keyWidth, descWidth, splitBody := splitHelp(body)
		w = keyWidth + descWidth + 2
		h := len(lines) + 1
		w = w + 1

		return &HelpWidget{
			BaseWidget: BaseWidget{
				name: name,
				x:    x,
				y:    y,
				w:    w,
				h:    h,
			},
			body:      body,
			splitBody: splitBody,
			keyWidth:  keyWidth,
			descWidth: descWidth,
		}
	})
}

func (w *HelpWidget) Layout(g *gocui.Gui) error {
	v, err := g.SetView(string(w.name), w.x, w.y, w.x+w.w, w.y+w.h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for _, line := range w.splitBody {
			fmt.Fprintf(v, "%-"+strconv.Itoa(w.keyWidth)+"s: %-"+strconv.Itoa(w.descWidth)+"s\n", line["key"], line["desc"])
		}
	}
	return nil
}

func splitHelp(body string) (int, int, []map[string]string) {
	lines := strings.Split(body, "\n")
	keyWidth := 0
	descWidth := 0
	var splitedArr []map[string]string
	runes := []rune{':', '：'}
	for _, l := range lines {
		str := lib.SplitStringByChars(runes, l)
		if len(str[0]) > keyWidth {
			keyWidth = len(str[0])
		}
		if len(str[1]) > descWidth {
			descWidth = len(str[1])
		}
		splitedArr = append(splitedArr, map[string]string{"key": str[0], "desc": str[1]})
	}
	return keyWidth, descWidth, splitedArr
}
