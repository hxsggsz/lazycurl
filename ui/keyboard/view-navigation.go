package keyboard

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var viewNavigation = map[any]func(g *gocui.Gui, v *gocui.View) error{
	gocui.KeyArrowUp:   scrollUp,
	'k':                scrollUp,
	gocui.KeyArrowDown: scrollDown,
	'j':                scrollDown,
}

func RegisterGlobalViewNavigation(g *gocui.Gui) error {
	targetViews := []string{views.RESPONSE, views.RESPONSE_HEADERS, views.LOGS, "method_modal", "help_modal"}

	for key, handler := range viewNavigation {
		for _, viewName := range targetViews {
			var err error

			switch k := key.(type) {
			case gocui.Key:
				err = g.SetKeybinding(viewName, k, gocui.ModNone, handler)
			case rune:
				err = g.SetKeybinding(viewName, k, gocui.ModNone, handler)
			}

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func scrollDown(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	cx, cy := v.Cursor()
	ox, oy := v.Origin()
	_, vy := v.Size()

	lineCount := len(v.BufferLines())

	if cy+oy >= lineCount-1 {
		return nil
	}

	if cy >= vy-1 {
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	} else {
		if err := v.SetCursor(cx, cy+1); err != nil {
			return err
		}
	}
	return nil
}

func scrollUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	cx, cy := v.Cursor()
	ox, oy := v.Origin()

	if cy > 0 {
		return v.SetCursor(cx, cy-1)
	}

	if oy > 3 {
		return v.SetOrigin(ox, oy-1)
	}

	return nil
}
