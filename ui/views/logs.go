package views

import (
	"github.com/awesome-gocui/gocui"
)

func Logs(g *gocui.Gui, maxX, maxY int) error {
	width, height := int(float64(maxX)*0.8), int(float64(maxY)*0.6)
	x0, y0 := (maxX-width)/2, (maxY-height)/2
	x1, y1 := x0+width, y0+height

	if v, err := g.SetView(LOGS, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Logs"
		v.Wrap = true
		v.Frame = false

		g.SetViewOnBottom(LOGS)

		if err := g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, toggleLogs); err != nil {
			return err
		}
	}
	return nil
}

var isLogsVisible = false

func toggleLogs(g *gocui.Gui, v *gocui.View) error {
	isLogsVisible = !isLogsVisible

	if isLogsVisible {
		views := g.Views()
		for _, view := range views {
			view.FrameColor = gocui.ColorWhite
			view.TitleColor = gocui.ColorWhite
		}

		g.SetViewOnTop(LOGS)
		v, err := g.SetCurrentView(LOGS)
		if err != nil {
			return err
		}

		v.Frame = true
		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen

		return nil
	}

	g.SetViewOnBottom(LOGS)
	v.Frame = false

	return nil
}
