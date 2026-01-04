package ui

import (
	"lazycurl/ui/keyboard"
	"lazycurl/ui/views"
	"lazycurl/ui/views/collection"
	"lazycurl/ui/views/helper"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := collection.Method(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Input(g, maxX); err != nil {
		return err
	}

	if err := collection.Headers(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Body(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Response(g, maxX, maxY); err != nil {
		return err
	}

	if err := views.Logs(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Help(g, maxX, maxY); err != nil {
		return err
	}

	keyboard.RegisterGlobalNumericNavigation(g)
	keyboard.RegisterGlobalSubmit(g)
	keyboard.RegisterGlobalViewNavigation(g)
	helper.ChangeViewFrame(g)

	return nil
}
