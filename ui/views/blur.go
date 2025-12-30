package views

import "github.com/awesome-gocui/gocui"

func blurInput(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView(LOGS)
	v.FrameColor = gocui.ColorWhite
	v.TitleColor = gocui.ColorWhite
	return err
}

func HandleBlurInput(g *gocui.Gui, viewName string) error {
	if err := g.SetKeybinding(viewName, gocui.KeyEsc, gocui.ModNone, blurInput); err != nil {
		return err
	}
	return nil
}
