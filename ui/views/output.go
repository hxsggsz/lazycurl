package views

import (
	"lazycurl/ui/utils"

	"github.com/awesome-gocui/gocui"
)

func Output(g *gocui.Gui, maxX, maxY int) error {
	if v, err := g.SetView(LOGS, 0, maxY-3, maxX-1, maxY-1, 0); err != nil {
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
