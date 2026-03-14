package helper

import (
	"github.com/awesome-gocui/gocui"
)

var (
	isViewVisible = false
)

func ToggleView(viewName string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		isViewVisible = !isViewVisible

		if isViewVisible {
			views := g.Views()
			for _, view := range views {
				view.FrameColor = gocui.ColorWhite
				view.TitleColor = gocui.ColorWhite
			}

			g.SetViewOnTop(viewName)
			v, err := g.SetCurrentView(viewName)
			if err != nil {
				return err
			}

			v.Visible = true

			v.Frame = true
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen

			return nil
		}

		v, err := g.SetViewOnBottom(viewName)
		if err != nil {
			return err
		}

		v.Visible = false

		return nil
	}
}
