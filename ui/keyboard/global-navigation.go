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

var focusableViews = []string{
	views.METHOD,
	views.URL,
	views.BODY,
	views.RESPONSE,
	views.HEADERS,
	views.RESPONSE_HEADERS,
	views.FILE_TREE_VIEW,
	views.ADD_FILE_URL,
	views.ADD_FILE_METHOD,
	views.ADD_FILE_NAME,
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

func RegisterGlobalMouseNavigation(g *gocui.Gui) error {
	return g.SetKeybinding("", gocui.MouseLeft, gocui.ModNone, handleGlobalMouseClick)
}

func handleGlobalMouseClick(g *gocui.Gui, v *gocui.View) error {
	log.Println("clicked")
	mx, my := g.MousePosition()

	clicked, err := g.ViewByPosition(mx, my)
	if err != nil {
		return nil
	}

	log.Println(clicked)

	name := clicked.Name()
	if name == "" {
		return nil
	}

	log.Printf("mouse click focused view: %s", name)

	for _, n := range focusableViews {
		if sv, err := g.View(n); err == nil {
			sv.FrameColor = gocui.ColorWhite
			sv.TitleColor = gocui.ColorWhite
		}
	}

	target, err := g.SetCurrentView(name)
	if err != nil {
		return err
	}
	target.FrameColor = gocui.ColorGreen
	target.TitleColor = gocui.ColorGreen

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
