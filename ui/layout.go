package ui

import (
	"lazycurl/ui/config"
	"lazycurl/ui/keyboard"
	"lazycurl/ui/views"
	"lazycurl/ui/views/collection"
	"lazycurl/ui/views/helper"

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

	if err := collection.FileTree(g, maxX, maxY, rootTree); err != nil {
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
	config.ShowCursor(g)

	return nil
}
