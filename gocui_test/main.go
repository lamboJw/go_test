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
	score := widget.GetScoreWidget()
	g.SetManager(help, next, main, score)
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
