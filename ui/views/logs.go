package views

import (
	"fmt"

	"lazycurl/ui/views/helper"

	"github.com/awesome-gocui/gocui"
)

func Logs(g *gocui.Gui, maxX, maxY int) error {
	viewWidth := 18
	x0, y0 := maxX-viewWidth-22, maxY-2
	x1, y1 := maxX-1, maxY

	if v, err := g.SetView(LOGS, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Label verde "F10 Logs"
		v.Frame = false
		fmt.Fprintf(v, "%spress F10 for Logs |%s", GREEN, RESET)
		g.SetViewOnTop(LOGS)

		// Keybinding F10 para toggle
		if err := g.SetKeybinding("", gocui.KeyF10, gocui.ModNone, helper.ToggleView("logs_modal")); err != nil {
			return err
		}
	}

	ShowLogsModal(g)
	return nil
}

func ShowLogsModal(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	w, h := 80, 25 // Maior para logs
	x0, y0 := (maxX-w)/2, (maxY-h)/2
	x1, y1 := x0+w, y0+h

	if v, err := g.SetView("logs_modal", x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		g.SetViewOnBottom("logs")
		v.Title = "Logs"
		v.FrameColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true
		v.Wrap = true       // Logs quebram linha
		v.Autoscroll = true // Scroll automático
		v.Visible = false

		if err := g.SetKeybinding("logs", gocui.KeyEsc, gocui.ModNone, helper.CloseView("logs")); err != nil {
			return err
		}
	}

	return nil
}
