package input

import "github.com/awesome-gocui/gocui"

func CreateInputView(g *gocui.Gui, maxX int) error {
	if v, err := g.SetView("input", 0, 0, maxX-1, 2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "URL"
		v.FrameColor = gocui.ColorGreen
		v.Editable = true
		_, _ = g.SetCurrentView("input")
	}
	return nil
}
