package gocui_test

import (
	"github.com/lamboJw/gocui"
	"go_test/gocui_test/lib"
	"go_test/gocui_test/widget"
	"log"
)

func Main() {
	g := lib.GetGui()
	defer g.Close()
	help := widget.GetHelpWidget()
	next := widget.GetNextWidget()
	main := widget.GetMainWidget()
	g.SetManager(help, next, main)
	g.Update(func(gui *gocui.Gui) error {
		setWidgetPos()
		return nil
	})
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		if main.ExistsCurDiamonds() {
			return nil
		}
		if err := main.InitWidgetLimitPos(); err != nil {
			log.Panicln(err)
		}
		err := main.RandDiamonds()
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		main.CurDiamondsMove(lib.DirectionLeft)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		main.CurDiamondsMove(lib.DirectionRight)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		main.CurDiamondsMove(lib.DirectionUp)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(gui *gocui.Gui, view *gocui.View) error {
		main.CurDiamondsMove(lib.DirectionDown)
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func setWidgetPos() {
	var widgetArr = []widget.GetSetter{widget.GetMainWidget(), widget.GetNextWidget()}
	for _, b := range widgetArr {
		left, right, top, bottom, midX, midY, err := lib.GetViewPos(b.Name())
		if err != nil {
			log.Panicln(err)
		}
		b.SetLeft(left)
		b.SetRight(right)
		b.SetTop(top)
		b.SetBottom(bottom)
		b.SetMidY(midY)
		b.SetMidX(midX)
	}
}
