package views

import (
	"fmt"
	"lazycurl/ui/utils"

	"github.com/awesome-gocui/gocui"
)

func Body(g *gocui.Gui, maxX, maxY int) error {
	height := maxY - 3 // reservar espaço para a view de logs

	x0 := 0
	x1 := maxX / 2

	y0 := 3
	y1 := height - 1

	if v, err := g.SetView(BODY, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf("[%d] %s", 2, utils.Capitalize(BODY))
		v.Autoscroll = true
		v.Editable = true
		v.Wrap = true
		g.Cursor = true
		g.SetViewOnBottom(BODY)
	}
	return nil
}
