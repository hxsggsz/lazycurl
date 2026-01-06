package views

import (
	"fmt"
	"lazycurl/pkg/highlight"
	"log"

	"github.com/awesome-gocui/gocui"
)

func blurInput(viewName string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		_, err := g.SetCurrentView(LOGS)
		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		if viewName == v.Name() {
			bodyInput := v.Buffer()
			v.Clear()

			log.Println("Blurring and highlighting JSON body")
			fmt.Fprint(v, highlight.Json(bodyInput))
		}

		return err
	}
}

func HandleBlurInput(g *gocui.Gui, viewName string) error {

	if err := g.SetKeybinding(viewName, gocui.KeyEsc, gocui.ModNone, blurInput(viewName)); err != nil {
		return err
	}
	return nil
}
