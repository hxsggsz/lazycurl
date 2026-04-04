package filetree

import (
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func addFile(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	width := 40
	height := 6
	x0 := (maxX - width) / 2
	y0 := (maxY - height) / 2
	x1 := x0 + width
	y1 := y0 + height

	v, err := g.SetView(views.ADD_FILE, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "New Request File"
		v.Frame = true
		v.FrameColor = gocui.ColorCyan
		v.TitleColor = gocui.ColorCyan
		v.Visible = false
	}

	requestNameInput(g, x0, x1, y0)
	methodSelectInput(g, x0, y0)
	requestNameInput(g, x0, x1, y0)

	g.SetKeybinding(views.ADD_FILE, gocui.KeyCtrlQ, gocui.ModNone, closeAddFileModal)
	g.SetKeybinding(views.ADD_FILE, gocui.KeyEsc, gocui.ModNone, closeAddFileModal)

	return nil
}

func toggleAddFileModal(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		outerV, err := g.View(views.ADD_FILE)
		if err != nil {
			return err
		}

		if outerV.Visible {
			outerV.Visible = false
			if nameV, err := g.View(views.ADD_FILE_NAME); err == nil {
				nameV.Visible = false
			}
			if methodV, err := g.View(views.ADD_FILE_METHOD); err == nil {
				methodV.Visible = false
			}
			g.DeleteView("add_file_method_modal")
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		outerV.Visible = true

		if methodV, err := g.View(views.ADD_FILE_METHOD); err == nil {
			methodV.Visible = true
			methodV.FrameColor = gocui.ColorGreen
			methodV.TitleColor = gocui.ColorGreen
		}

		if nameV, err := g.View(views.ADD_FILE_NAME); err == nil {
			nameV.Clear()
			nameV.Visible = true
			nameV.FrameColor = gocui.ColorGreen
			nameV.TitleColor = gocui.ColorGreen
		}

		g.SetViewOnTop(views.ADD_FILE)
		g.SetViewOnTop(views.ADD_FILE_NAME)
		g.SetViewOnTop(views.ADD_FILE_METHOD)
		g.SetCurrentView(views.ADD_FILE_METHOD)

		return nil
	}
}

func closeAddFileModal(g *gocui.Gui, v *gocui.View) error {
	if outerV, err := g.View(views.ADD_FILE); err == nil {
		outerV.Visible = false
	}
	if nameV, err := g.View(views.ADD_FILE_NAME); err == nil {
		nameV.Visible = false
	}
	if methodV, err := g.View(views.ADD_FILE_METHOD); err == nil {
		methodV.Visible = false
	}
	g.DeleteView("add_file_method_modal")
	g.SetCurrentView(views.FILE_TREE_VIEW)
	return nil
}
