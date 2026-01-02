package collection

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func Headers(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.HEADERS

	height := maxY - views.LOGS_HEIGHT

	x0 := views.FULL
	x1 := maxX / 2

	y0 := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
	y1 := height - views.LOGS_BOTTOM

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = true
		v.Wrap = true
		v.Title = "[3]  Body *Headers"

		if err := g.SetKeybinding(viewName, gocui.KeyArrowLeft, gocui.ModShift, prevTab); err != nil {
			return err
		}
		if err := g.SetKeybinding(viewName, gocui.KeyArrowRight, gocui.ModShift, nextTab); err != nil {
			return err
		}
	}

	return nil
}
