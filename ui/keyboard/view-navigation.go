package keyboard

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var viewNavigation = map[gocui.Key]func(g *gocui.Gui, v *gocui.View) error{
	gocui.KeyArrowUp: scrollUp, 'k': scrollUp,
	gocui.KeyArrowDown: scrollDown, 'j': scrollDown,
}

func RegisterGlobalViewNavigation(g *gocui.Gui) error {
	for key, handler := range viewNavigation {
		if err := g.SetKeybinding(views.RESPONSE, key, gocui.ModNone, handler); err != nil {
			return err
		}
	}
	return nil
}

func scrollDown(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		if v != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func scrollUp(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		if v != nil {
			ox, oy := v.Origin()
			if oy > 0 {
				if err := v.SetOrigin(ox, oy-1); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return nil
}
