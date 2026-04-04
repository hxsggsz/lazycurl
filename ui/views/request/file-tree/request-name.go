package filetree

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func requestNameInput(g *gocui.Gui, x0, x1, y0 int) error {
	nameV, err := g.SetView(views.ADD_FILE_NAME, x0+12, y0+5, x1-2, y0+7, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		nameV.Editable = true
		nameV.Title = "Request name:"
		nameV.FrameColor = gocui.ColorWhite
		nameV.TitleColor = gocui.ColorWhite
		nameV.Visible = false
	}

	return nil
}
