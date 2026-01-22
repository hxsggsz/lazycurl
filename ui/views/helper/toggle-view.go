package helper

import "github.com/awesome-gocui/gocui"

var (
	isLogsVisible = false
)

func CloseView(viewName string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		isLogsVisible = false

		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite
		v.Frame = false

		g.SetViewOnBottom(viewName)
		g.Cursor = true

		return nil
	}
}

func ToggleView(viewName string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		isLogsVisible = !isLogsVisible

		if isLogsVisible {
			views := g.Views()
			g.Cursor = false
			for _, view := range views {
				view.FrameColor = gocui.ColorWhite
				view.TitleColor = gocui.ColorWhite
			}

			g.SetViewOnTop(viewName)
			v, err := g.SetCurrentView(viewName)
			if err != nil {
				return err
			}

			v.Frame = true
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen

			return nil
		}

		g.Cursor = true
		g.SetViewOnBottom(viewName)
		v.Frame = false

		return nil
	}
}
