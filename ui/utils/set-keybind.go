package utils

import (
	"github.com/awesome-gocui/gocui"
)

type keybindsMaps map[gocui.Key]func(g *gocui.Gui, v *gocui.View) error

func SetKeybind(g *gocui.Gui, kbm keybindsMaps, viewName string) error {
	for key, handler := range kbm {
		if err := g.SetKeybinding(viewName, key, gocui.ModNone, handler); err != nil {
			return err
		}
	}

	return nil
}
