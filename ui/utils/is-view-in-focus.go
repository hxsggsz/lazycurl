package utils

import (
	"github.com/awesome-gocui/gocui"
)

func ViewHasFocus(g *gocui.Gui, viewName string) bool {
	thisView, _ := g.View(viewName)
	return g.CurrentView() == thisView
}
