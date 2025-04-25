package ui

import (
	"lazycurl/ui/components/input"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := input.CreateInputView(g, maxX); err != nil {
		return err
	}

	if _, err := g.SetView("main", 0, 3, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}
