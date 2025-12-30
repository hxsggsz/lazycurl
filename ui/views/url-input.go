package views

import (
	"fmt"
	"lazycurl/ui/utils"
	"log"

	"github.com/awesome-gocui/gocui"
)

func inputViewConfig(g *gocui.Gui, v *gocui.View) {
	v.Title = fmt.Sprintf("[%d] %s", 1, utils.Capitalize(URL))
	v.Editable = true
	g.Cursor = true

}

func blurInput(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(LOGS)
	v.FrameColor = gocui.ColorWhite
	v.TitleColor = gocui.ColorWhite
	return err
}

func Input(g *gocui.Gui, maxX int) (string, error) {
	viewName := URL

	if v, err := g.SetView(viewName, 0, 0, maxX-1, 2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return "", err
		}

		// default focus when iniciate the view
		g.SetCurrentView(viewName)
		inputViewConfig(g, v)

		if err := g.SetKeybinding(viewName, gocui.KeyEsc, gocui.ModNone, blurInput); err != nil {
			return "", err
		}

		if hasFocus := utils.ViewHasFocus(g, viewName); hasFocus {
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen
		}
	}

	inputView, err := g.View(viewName)
	if err != nil {
		log.Panicf("not found view -> %v", viewName)
	}

	return inputView.Buffer(), nil
}
