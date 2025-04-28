package components

import (
	"lazycurl/ui/utils"
	"log"

	"github.com/awesome-gocui/gocui"
)

func inputViewConfig(g *gocui.Gui, v *gocui.View, viewName string) {
	v.Title = "URL"
	v.Editable = true
	g.Cursor = true

	if hasFocus := utils.ViewHasFocus(g, viewName); hasFocus {
		v.FrameColor = gocui.ColorGreen
	}
}

func Input(g *gocui.Gui, maxX int) (string, error) {
	viewName := "input"
	if v, err := g.SetView(viewName, 0, 0, maxX-1, 2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return "", err
		}

		inputViewConfig(g, v, viewName)

		// default focus when iniciate the view
		g.SetCurrentView(viewName)

	}

	inputView, err := g.View(viewName)
	if err != nil {
		log.Panicf("not found view -> %v", viewName)
	}

	return inputView.Buffer(), nil
}
