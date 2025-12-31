package keyboard

import (
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

var globalViewMap = map[rune]string{
	'1': views.URL, '2': views.BODY,
	'3': views.RESPONSE,
}

func RegisterGlobalNumericNavigation(g *gocui.Gui) error {
	for key, viewName := range globalViewMap {
		if err := g.SetKeybinding("", key, gocui.ModNone, makeFocusHandler(viewName)); err != nil {
			return err
		}
	}
	return nil
}

func makeFocusHandler(viewName string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		log.Printf("changing focus to section: %s", viewName)
		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		newView, err := g.SetCurrentView(viewName)
		newView.FrameColor = gocui.ColorGreen
		newView.TitleColor = gocui.ColorGreen

		return err
	}
}
