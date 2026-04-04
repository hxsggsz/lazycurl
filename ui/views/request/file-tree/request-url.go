package filetree

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func requestUrlInput(g *gocui.Gui, x0, x1, y0 int) error {
	urlV, err := g.SetView(views.ADD_FILE_URL, x0+2, y0+5, x1-2, y0+7, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		urlV.Editable = true
		urlV.Title = "Request url:"
		urlV.FrameColor = gocui.ColorWhite
		urlV.TitleColor = gocui.ColorWhite
		urlV.Visible = false
	}

	return nil
}
