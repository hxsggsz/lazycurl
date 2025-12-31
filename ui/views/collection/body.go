package collection

import (
	"fmt"
	"lazycurl/ui/utils"
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

		v.Title = fmt.Sprintf("[%d] %s", 3, utils.Capitalize(viewName))
		v.Autoscroll = true
		v.Editable = true
		v.Wrap = true
		g.Cursor = true
		g.SetViewOnBottom(viewName)
		views.HandleBlurInput(g, viewName)
	}
	return nil
}
