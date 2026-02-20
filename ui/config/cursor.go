package config

import "github.com/awesome-gocui/gocui"

func ShowCursor(g *gocui.Gui) {
	currentView := g.CurrentView()
	if isEditable := currentView != nil && currentView.Editable; isEditable {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

}
