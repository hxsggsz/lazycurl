package utils

import (
	"github.com/awesome-gocui/gocui"
)

// KeybindAction agrupa o modificador e a função de execução
type KeybindAction struct {
	Modifier gocui.Modifier
	Handler  func(g *gocui.Gui, v *gocui.View) error
}

// KeybindsMaps agora mapeia a tecla para a nossa estrutura de ação
type KeybindsMaps map[gocui.Key]KeybindAction

func SetKeybind(g *gocui.Gui, kbm KeybindsMaps, viewName string) error {
	for key, action := range kbm {
		if err := g.SetKeybinding(viewName, key, action.Modifier, action.Handler); err != nil {
			return err
		}
	}

	return nil
}
