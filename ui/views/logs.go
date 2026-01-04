package views

import (
	"lazycurl/ui/views/helper"

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
		g.Cursor = false

		v.Title = "Logs"
		v.Wrap = true
		v.Frame = false

		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen

		g.SetViewOnBottom(LOGS)

		if err := g.SetKeybinding(LOGS, gocui.KeyF10, gocui.ModNone, helper.ToggleLogs(LOGS)); err != nil {
			return err
		}

		if err := g.SetKeybinding(LOGS, gocui.KeyEsc, gocui.ModNone, helper.ToggleLogs(LOGS)); err != nil {
			return err
		}
	}
	return nil
}
