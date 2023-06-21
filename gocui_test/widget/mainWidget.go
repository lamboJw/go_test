package widget

import (
	"fmt"
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/diamonds"
	"go_test/gocui_test/lib"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

type MainWidget struct {
	BaseWidget
	diamondsNum      int
	curDiamonds      diamonds.Eventer
	level            int
	score            int
	existsDiamond    map[int]map[int]*diamonds.Diamond
	bottomChannel    chan bool
	topChannel       chan bool
	directionChannel chan lib.Direction
	stopChannel      chan bool
	lock             sync.RWMutex
}

func init() {
	creatorRegister(lib.MainWidgetName, NewMainWidget)
}

func NewMainWidget(name lib.WidgetName, x int, y int, args ...interface{}) Widget {
	w := args[0].(int)
	h := args[1].(int)
	if w%lib.DiamondWidth != 0 {
		log.Panicln("宽度不是方块宽度的整数倍")
	}
	if h%lib.DiamondHeight != 0 {
		log.Panicln("高度不是方块高度的整数倍")
	}
	d := &MainWidget{
		BaseWidget: BaseWidget{
			name: name,
			x:    x,
			y:    y,
			w:    w + 1,
			h:    h + 1,
		},
		diamondsNum:      0,
		level:            1,
		score:            0,
		curDiamonds:      nil,
		bottomChannel:    make(chan bool, 1),
		topChannel:       make(chan bool, 1),
		stopChannel:      make(chan bool, 1),
		directionChannel: make(chan lib.Direction, 10),
	}
	return d
}

func (w *MainWidget) Layout(g *gocui.Gui) error {
	_, err := g.SetView(string(w.name), w.x, w.y, w.x+w.w+1, w.y+w.h+1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

func (w *MainWidget) AddFixedDiamond() {
	for _, diamond := range w.curDiamonds.GetDiamondArr() {
		pos := diamond.GetPos()
		w.existsDiamond[pos[1]][pos[0]] = diamond
	}
	w.drawExistsDiamond()
	w.curDiamonds = nil
}

func (w *MainWidget) SetCurDiamonds(d diamonds.Eventer) {
	w.curDiamonds = d
	w.startEvent()
}

func (w *MainWidget) startEvent() {
	go func() {
		g := lib.GetGui()
		for {
			select {
			case <-w.stopChannel:
				return
			case <-time.After(w.getSpeed()):
				g.Update(func(gui *gocui.Gui) error {
					if !w.checkStop(lib.DirectionDown) {
						if err := w.curDiamonds.MoveDown(); err != nil {
							return err
						}
					}
					return nil
				})
			}
		}
	}()
	go func() {
		g := lib.GetGui()
		for {
			select {
			case <-w.bottomChannel:
				w.AddFixedDiamond()
				w.eliminate()
				g.Update(func(gui *gocui.Gui) error {
					if err := w.RandDiamonds(); err != nil {
						return err
					}
					return nil
				})
				return
			case <-w.topChannel:
				w.curDiamonds = nil
				g.Update(func(gui *gocui.Gui) error {
					maxX, maxY := g.Size()
					v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2)
					if err != nil && err != gocui.ErrUnknownView {
						return err
					}
					fmt.Fprintln(v, "游戏结束")
					g.Highlight = true
					g.SelFgColor = gocui.ColorRed
					g.SelBgColor = gocui.ColorGreen
					g.SetCurrentView("msg")
					if _, err := g.SetViewOnTop("msg"); err != nil {
						return err
					}
					return nil
				})
				return
			case direction := <-w.directionChannel:
				if w.ExistsCurDiamonds() {
					switch direction {
					case lib.DirectionLeft:
						g.Update(func(gui *gocui.Gui) error {
							if !w.checkStop(direction) {
								if err := w.curDiamonds.MoveLeft(); err != nil {
									return err
								}
							}
							return nil
						})
					case lib.DirectionRight:
						g.Update(func(gui *gocui.Gui) error {
							if !w.checkStop(direction) {
								if err := w.curDiamonds.MoveRight(); err != nil {
									return err
								}
							}
							return nil
						})
					case lib.DirectionDown:
						g.Update(func(gui *gocui.Gui) error {
							for !w.checkStop(direction) {
								if err := w.curDiamonds.MoveDown(); err != nil {
									return err
								}
							}
							return nil
						})
					case lib.DirectionUp:
						g.Update(func(gui *gocui.Gui) error {
							diamondArr, switchType := w.curDiamonds.GetSwitchDirectionPos()
							canSwitch := true
							for _, pos := range diamondArr {
								xMap := w.existsDiamond[pos[1]]
								diamond, ok := xMap[pos[0]]
								if !ok {
									canSwitch = false
									break
								}
								if diamond != nil {
									canSwitch = false
									break
								}
							}
							if canSwitch {
								if err := w.curDiamonds.SwitchDirection(diamondArr, switchType); err != nil {
									return err
								}
							}
							return nil
						})
					}
				}
			}
		}
	}()
}

func (w *MainWidget) getSpeed() time.Duration {
	speed := 11 - w.level
	if speed < 1 {
		speed = 1
	}
	return time.Duration(speed) * 100 * time.Millisecond
}

func (w *MainWidget) CurDiamondsMove(direction lib.Direction) {
	w.lock.Lock()
	w.directionChannel <- direction
	w.lock.Unlock()
}

func (w *MainWidget) ExistsCurDiamonds() bool {
	return w.curDiamonds != nil
}

func (w *MainWidget) getInitExistsDiamond() map[int]map[int]*diamonds.Diamond {
	existsDiamond := make(map[int]map[int]*diamonds.Diamond)
	for y := w.top - 5*lib.DiamondHeight; y < w.bottom; y += lib.DiamondHeight { // 上面加高5格，方便在方块刚生成时判断是否可以变换方向
		existsDiamond[y] = make(map[int]*diamonds.Diamond)
		for x := w.left; x < w.right; x += lib.DiamondWidth {
			existsDiamond[y][x] = nil
		}
	}
	return existsDiamond
}

func (w *MainWidget) InitWidgetLimitPos() error {
	if w.existsDiamond != nil {
		for _, xArray := range w.existsDiamond {
			for _, diamond := range xArray {
				if diamond != nil {
					diamond.Destroy()
				}
			}
		}
	}
	w.existsDiamond = w.getInitExistsDiamond()
	w.drawExistsDiamond()
	return nil
}

//var existsDiamondCount = 0

func (w *MainWidget) drawExistsDiamond() {
	//existsDiamondCount++
	fp, _ := os.OpenFile("existsDiamond.txt", os.O_WRONLY|os.O_CREATE, 0644)
	defer fp.Close()
	for y := w.top - 5*lib.DiamondHeight; y < w.bottom; y += lib.DiamondHeight {
		for x := w.left; x < w.right; x += lib.DiamondWidth {
			if w.existsDiamond[y][x] != nil {
				fp.WriteString("1 ")
			} else {
				fp.WriteString("0 ")
			}
		}
		fp.WriteString("\n")
		if y == w.top {
			for x := w.left; x < w.right; x += lib.DiamondWidth {
				fp.WriteString("- ")
			}
			fp.WriteString("\n")
		}
	}
}

func (w *MainWidget) checkStop(direction lib.Direction) bool {
	result := false
	var bottomY = -9999
	diamondArr := w.curDiamonds.GetDiamondArr()
	for _, diamond := range diamondArr {
		pos := diamond.GetPos()
		x := pos[0]
		y := pos[1]
		if y > bottomY {
			bottomY = y
		}
		switch direction {
		case lib.DirectionLeft:
			x -= lib.DiamondWidth
		case lib.DirectionRight:
			x += lib.DiamondWidth
		case lib.DirectionDown:
			y += lib.DiamondHeight
		}
		if x < w.left || x > w.right-lib.DiamondWidth || y > w.bottom-lib.DiamondHeight {
			result = true
		} else if y >= w.top && w.existsDiamond[y][x] != nil {
			result = true
		}
	}
	if result && direction == lib.DirectionDown {
		w.lock.Lock()
		if bottomY < w.top {
			w.topChannel <- true
		} else {
			w.bottomChannel <- true
		}
		w.stopChannel <- true
		w.lock.Unlock()
	}
	return result
}

func (w *MainWidget) RandDiamonds() error {
	randType := diamonds.GetRandomType()
	var d, d1 diamonds.Eventer
	var err error
	if err = next.DestroyNextDiamonds(); err != nil {
		return err
	}
	if d, err = diamonds.Create(randType, -1, lib.NextWidgetName); err != nil {
		return err
	}
	if err = next.SetNextDiamonds(d); err != nil {
		return err
	}
	w.diamondsNum++
	if d1, err = diamonds.Create(randType, w.diamondsNum, lib.MainWidgetName); err != nil {
		return err
	}
	main.SetCurDiamonds(d1)
	return nil
}

func (w *MainWidget) eliminate() {
	eliminateY := make(map[int]bool, w.w/lib.DiamondWidth)
	eliminateYArr := make(sort.IntSlice, 0)
	for y, xArray := range w.existsDiamond {
		if y < w.top {
			continue
		}
		full := true
		for _, diamond := range xArray {
			if diamond == nil {
				full = false
				break
			}
		}
		if full {
			eliminateY[y] = true
			eliminateYArr = append(eliminateYArr, y)
		}
	}
	if len(eliminateY) == 0 {
		return
	}
	eliminateYArr.Sort()
	lib.GetGui().Update(func(gui *gocui.Gui) error {
		for index := len(eliminateYArr) - 1; index >= 0; index-- {
			for y, xArray := range w.existsDiamond {
				if y < w.top {
					continue
				}
				if y == eliminateYArr[index] { // 删除行
					for x, diamond := range xArray {
						diamond.Destroy()
						w.existsDiamond[y][x] = nil
					}
					continue
				}
				if _, ok := eliminateY[y]; ok || y > eliminateYArr[index] { // 大于当前要删除的行，或者这一行本来就要删除，什么都不执行
					continue
				}
				for _, diamond := range xArray {
					if diamond != nil {
						diamond.MoveDown()
					}
				}
			}
		}
		newExistsDiamond := w.getInitExistsDiamond() // 生成新的矩阵
		for y, xArray := range w.existsDiamond {
			if y < w.top {
				continue
			}
			for _, diamond := range xArray {
				if diamond != nil {
					pos := diamond.GetPos()
					newExistsDiamond[pos[1]][pos[0]] = diamond //把方块的新坐标对应到新的矩阵
				}
			}
		}
		w.existsDiamond = newExistsDiamond
		w.drawExistsDiamond()
		w.score += len(eliminateYArr)
		w.level = w.score/10 + 1
		log.Println(w.score, w.level)
		return nil
	})
}
