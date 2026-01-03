package collection

import (
	"fmt"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func Response(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.RESPONSE
	height := maxY - views.LOGS_HEIGHT

	width := maxX / 2
	x0 := width + views.LAYOUT_SECTION_X_GAP
	x1 := maxX - views.RIGHT_BORDER

	y0 := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
	y1 := height - views.LOGS_BOTTOM

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf("[%d] %s", 4, utils.Capitalize(viewName))
		v.Autoscroll = false
		v.Wrap = true
		v.HasLoader = true
		g.Cursor = true
	}

	return nil
}

