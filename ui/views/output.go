package views

import (
	"lazycurl/ui/utils"

	"github.com/awesome-gocui/gocui"
)

func Output(g *gocui.Gui, maxX, maxY int) error {

	x0 := FULL
	y0 := maxY - LOGS_HEIGHT

	x1 := maxX - LAYOUT_SECTION_Y_GAP
	y1 := maxY - RIGHT_BORDER

	if v, err := g.SetView(LOGS, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = utils.Capitalize(LOGS)
		v.Autoscroll = true
		v.Wrap = true
		g.SetViewOnBottom(LOGS)
	}

	return nil
}
