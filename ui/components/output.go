package components

import "github.com/awesome-gocui/gocui"

func Output(g *gocui.Gui, maxX, maxY int) error {
	if v, err := g.SetView("logs", 0, maxY-5, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Logs"
		v.Autoscroll = true
		v.Wrap = true
		g.SetViewOnBottom("logs")
	}
	return nil
}
