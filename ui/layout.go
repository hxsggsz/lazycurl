package ui

import (
	"lazycurl/pkg/collection"
	"lazycurl/ui/config"
	"lazycurl/ui/keyboard"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"lazycurl/ui/views/request"
	filetree "lazycurl/ui/views/request/file-tree"

	"github.com/awesome-gocui/gocui"
)

var rootTree = []collection.FileNode{
	{
		Name:  "root",
		IsDir: true,
		Open:  true,
		Children: []collection.FileNode{
			{Name: "arquivo 1", IsDir: false},
			{
				Name:  "user",
				IsDir: true,
				Open:  true,
				Children: []collection.FileNode{
					{Name: "arquivo 2", IsDir: false},
					{Name: "arquivo 3", IsDir: false},
					{
						Name:  "payment",
						IsDir: true,
						Open:  true,
						Children: []collection.FileNode{
							{Name: "arquivo 4", IsDir: false},
							{Name: "arquivo 5", IsDir: false},
						},
					},
				},
			},
		},
	},
	{
		Name:  "consumer",
		IsDir: true,
		Open:  false,
		Children: []collection.FileNode{
			{Name: "nome da pasta", IsDir: true},
		},
	},
}

func layout(collection *collection.Collection) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		maxX, maxY := g.Size()

		if err := request.Method(g, maxX, maxY); err != nil {
			return err
		}

		if err := request.Input(g, maxX); err != nil {
			return err
		}

		isMenuOpen, err := filetree.FileTree(g, maxX, maxY, collection.Files, false, collection.AddFolders)
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
		helper.ChangeViewFrame(g)
		config.ShowCursor(g)

		return nil
	}
}
