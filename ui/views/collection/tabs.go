package collection

import (
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var (
	BodyTabs     = []string{views.BODY, "header_key_0"}
	ResponseTabs = []string{views.RESPONSE, views.RESPONSE_HEADERS}
	ActiveTabIdx = 0
)

func nextTab(tabs []string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		ActiveTabIdx = (ActiveTabIdx + 1) % len(tabs)

		g.Update(func(g *gocui.Gui) error {
			for _, tabView := range tabs {
				g.SetViewOnBottom(tabView)
			}

			v.FrameColor = gocui.ColorWhite
			v.TitleColor = gocui.ColorWhite

			activeTab := tabs[ActiveTabIdx]
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
}

func prevTab(tabs []string) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		ActiveTabIdx--
		if ActiveTabIdx < 0 {
			ActiveTabIdx = len(tabs) - 1
		}

		g.Update(func(g *gocui.Gui) error {
			for _, tabView := range tabs {
				g.SetViewOnBottom(tabView)
			}

			v.FrameColor = gocui.ColorWhite
			v.TitleColor = gocui.ColorWhite

			activeTab := tabs[ActiveTabIdx]
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
}
