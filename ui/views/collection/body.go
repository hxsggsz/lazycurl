package collection

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func Body(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.BODY
	height := maxY - views.LOGS_HEIGHT

	x0 := views.FULL
	x1 := maxX / 2

	y0 := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
	y1 := height - views.LOGS_BOTTOM

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "[3] *Body  Headers"
		v.Editable = true
		v.Wrap = true
		v.Autoscroll = false

		g.Cursor = true
		views.HandleBlurInput(g, viewName)

		if err := g.SetKeybinding(viewName, gocui.KeyArrowLeft, gocui.ModShift, prevTab(BodyTabs)); err != nil {
			return err
		}
		if err := g.SetKeybinding(viewName, gocui.KeyArrowRight, gocui.ModShift, nextTab(BodyTabs)); err != nil {
			return err
		}
	}

	return nil
}

func GetBodyValue(g *gocui.Gui) string {
	v, err := g.View(views.BODY)
	if err != nil {
		return ""
	}

	return v.Buffer()
}
