package options

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func QuitKeyByind(g *gocui.Gui) {
	keysToQuit := []any{gocui.KeyCtrlC}

	for _, key := range keysToQuit {
		if err := g.SetKeybinding("", key, gocui.ModNone, quit); err != nil {
			log.Panicln(err)
		}
	}
}
