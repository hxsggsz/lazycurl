package ui

import (
	"lazycurl/pkg/collection"
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/config"
	"lazycurl/ui/keyboard"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"lazycurl/ui/views/request"
	filetree "lazycurl/ui/views/request/file-tree"

	"github.com/awesome-gocui/gocui"
)

func layout(collection *collection.Collection) func(g *gocui.Gui) error {
	fileManager := &fm.FileManager{Collection: collection}

	return func(g *gocui.Gui) error {
		maxX, maxY := g.Size()

		if err := request.Method(g, maxX, maxY); err != nil {
			return err
		}

		if err := request.Input(g, maxX); err != nil {
			return err
		}

		isMenuOpen, err := filetree.FileTree(g, maxX, maxY, fileManager, false)
		if err != nil {
			return err
		}

		if err := request.Headers(g, maxX, maxY, isMenuOpen); err != nil {
			return err
		}

		if err := request.Body(g, maxX, maxY, isMenuOpen); err != nil {
			return err
		}

		if err := request.Response(g, maxX, maxY); err != nil {
			return err
		}

		if err := views.Logs(g, maxX, maxY); err != nil {
			return err
		}

		if err := request.Help(g, maxX, maxY); err != nil {
			return err
		}

		keyboard.RegisterGlobalNumericNavigation(g)
		keyboard.RegisterGlobalSubmit(g)
		keyboard.RegisterGlobalViewNavigation(g)
		keyboard.RegisterGlobalMouseNavigation(g)
		helper.ChangeViewFrame(g)
		config.ShowCursor(g)

		return nil
	}
}
