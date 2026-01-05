package utils

import (
	"github.com/awesome-gocui/gocui"
)

type KeyCombo struct {
	Key      gocui.Key
	Modifier gocui.Modifier
}

type KeybindsMaps map[KeyCombo]func(g *gocui.Gui, v *gocui.View) error

func SetKeybind(g *gocui.Gui, kbm KeybindsMaps, viewName string) error {
	for kc, handler := range kbm {
		if err := g.SetKeybinding(viewName, kc.Key, kc.Modifier, handler); err != nil {
			return err
		}
	}

	return nil
}
