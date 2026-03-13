package keyboard

import (
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

var numericMap = map[rune]string{
	'1': views.METHOD,
	'2': views.URL,
	'3': views.BODY,
	'4': views.RESPONSE,
}

var specialKeyMap = map[gocui.Key]string{
	gocui.KeyTab: views.FILE_TREE_VIEW,
}

func RegisterGlobalNumericNavigation(g *gocui.Gui) error {
	for key, viewName := range numericMap {
		if err := g.SetKeybinding("", key, gocui.ModNone, makeFocusHandler(viewName)); err != nil {
			return err
		}
	}

	for key, viewName := range specialKeyMap {
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
