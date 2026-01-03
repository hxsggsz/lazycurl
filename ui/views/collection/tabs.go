package collection

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var TabNames = []string{views.BODY, "header_key_0"}
var ActiveTabIdx = 0

func nextTab(g *gocui.Gui, v *gocui.View) error {
	ActiveTabIdx = (ActiveTabIdx + 1) % len(TabNames)

	g.Update(func(g *gocui.Gui) error {
		for _, tabView := range TabNames {
			g.SetViewOnBottom(tabView)
		}

		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		activeTab := TabNames[ActiveTabIdx]
		g.SetViewOnTop(activeTab)
		if newView, err := g.SetCurrentView(activeTab); err != nil {
			return err
		} else {
			newView.FrameColor = gocui.ColorGreen
			newView.TitleColor = gocui.ColorGreen
		}

		return nil
	})

	return nil
}

func prevTab(g *gocui.Gui, v *gocui.View) error {
	ActiveTabIdx--
	if ActiveTabIdx < 0 {
		ActiveTabIdx = len(TabNames) - 1
	}

	g.Update(func(g *gocui.Gui) error {
		for _, tabView := range TabNames {
			g.SetViewOnBottom(tabView)
		}

		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		activeTab := TabNames[ActiveTabIdx]
		if _, err := g.SetViewOnTop(activeTab); err != nil {
			return err
		}

		if newView, err := g.SetCurrentView(activeTab); err != nil {
			return err
		} else {
			newView.FrameColor = gocui.ColorGreen
			newView.TitleColor = gocui.ColorGreen
		}

		return nil
	})

	return nil
}
