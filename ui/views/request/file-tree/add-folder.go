package filetree

import (
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
	"time"

	"github.com/awesome-gocui/gocui"
)

// addFolder creates the modal view for entering new folder names.
func addFolder(g *gocui.Gui, maxX, maxY int, fm *fm.FileManager) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView(views.ADD_FOLDER, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		v.Title = " Folder(s) Name(s) "
		v.Subtitle = " Separate multiple folders with slashes (e.g., folder-1/folder-2/) "
		v.Frame = true
		v.Wrap = true
		v.Editable = true
		v.Visible = false

		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen
	}

	g.SetKeybinding(views.ADD_FOLDER, gocui.KeyEnter, gocui.ModNone, createFolders(fm))
	g.SetKeybinding(views.ADD_FOLDER, gocui.KeyCtrlQ, gocui.ModNone, helper.CloseView(views.ADD_FOLDER))
	g.SetKeybinding(views.ADD_FOLDER, gocui.KeyEsc, gocui.ModNone, helper.CloseView(views.ADD_FOLDER))
	return nil
}

// getFoldersInput retrieves the current text from the add folder modal.
func getFoldersInput(g *gocui.Gui) string {
	v, err := g.View(views.ADD_FOLDER)
	if err != nil {
		log.Println("Erro ao obter view do modal:", err)
		return ""
	}
	return v.Buffer()
}

// clearFoldersInput clears the text in the add folder modal.
func clearFoldersInput(g *gocui.Gui) {
	if v, err := g.View(views.ADD_FOLDER); err == nil {
		v.Clear()
	}
}

// createFolders returns a keybinding handler that creates the entered folder
// structure, closes the modal, shows a success toast, and updates the tree view.
func createFolders(fm *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newFoldersPath := getFoldersInput(g)
		if newFoldersPath == "" {
			log.Println("folder name cannot be empty")
			return nil
		}

		fm.Collection.AddFolders(newFoldersPath)
		clearFoldersInput(g)
		helper.CloseView(views.ADD_FOLDER)(g, v)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		views.ShowToast(g, "folder created successfully", "success", 2*time.Second)
		fm.UpdateTree(g)
		return nil
	}
}
