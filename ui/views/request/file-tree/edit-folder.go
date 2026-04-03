package filetree

import (
	"fmt"
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
	"path/filepath"
	"time"

	"github.com/awesome-gocui/gocui"
)

// editFolder creates the modal view for renaming a folder.
func editFolder(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView(views.EDIT_FOLDER, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		v.Title = " Rename Folder "
		v.Frame = true
		v.Wrap = true
		v.Editable = true
		v.Visible = false

		v.FrameColor = gocui.ColorYellow
		v.TitleColor = gocui.ColorYellow
	}

	g.SetKeybinding(views.EDIT_FOLDER, gocui.KeyEnter, gocui.ModNone, renameFolder(fileManager))
	g.SetKeybinding(views.EDIT_FOLDER, gocui.KeyCtrlQ, gocui.ModNone, helper.CloseView(views.EDIT_FOLDER))
	g.SetKeybinding(views.EDIT_FOLDER, gocui.KeyEsc, gocui.ModNone, helper.CloseView(views.EDIT_FOLDER))
	return nil
}

// getEditFolderInput retrieves the current text from the edit folder modal.
func getEditFolderInput(g *gocui.Gui) string {
	v, err := g.View(views.EDIT_FOLDER)
	if err != nil {
		log.Println("Erro ao obter view do modal:", err)
		return ""
	}
	return v.Buffer()
}

// clearEditFolderInput clears the text in the edit folder modal.
func clearEditFolderInput(g *gocui.Gui) {
	if v, err := g.View(views.EDIT_FOLDER); err == nil {
		v.Clear()
	}
}

// populateEditFolderInput sets the initial folder name in the edit modal.
func populateEditFolderInput(g *gocui.Gui, name string) {
	if v, err := g.View(views.EDIT_FOLDER); err == nil {
		v.Clear()
		fmt.Fprint(v, name)
	}
}

// renameFolder returns a keybinding handler that renames the selected folder.
func renameFolder(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newName := getEditFolderInput(g)
		if newName == "" {
			log.Println("folder name cannot be empty")
			return nil
		}

		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FOLDER)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		if !node.IsDir {
			views.ShowToast(g, "selected item is not a folder", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FOLDER)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		rootPath := fileManager.Collection.GetRootPath()
		relPath, err := filepath.Rel(rootPath, node.Path)
		if err != nil {
			log.Printf("failed to get relative path: %v", err)
			views.ShowToast(g, "failed to rename folder", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FOLDER)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		if err := fileManager.Collection.RenameNode(relPath, newName); err != nil {
			log.Printf("failed to rename folder: %v", err)
			views.ShowToast(g, fmt.Sprintf("failed to rename: %v", err), "error", 2*time.Second)
			helper.CloseView(views.EDIT_FOLDER)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		clearEditFolderInput(g)
		helper.CloseView(views.EDIT_FOLDER)(g, v)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		views.ShowToast(g, fmt.Sprintf("%q renamed to %q", node.Name, newName), "success", 2*time.Second)
		fileManager.UpdateTree(g)
		return nil
	}
}

// triggerEditFolder opens the edit folder modal with the current folder name.
func triggerEditFolder(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			return nil
		}

		if !node.IsDir {
			views.ShowToast(g, "selected item is not a folder", "error", 2*time.Second)
			return nil
		}

		populateEditFolderInput(g, node.Name)

		if modalView, err := g.View(views.EDIT_FOLDER); err == nil {
			modalView.Visible = true
		}

		g.SetCurrentView(views.EDIT_FOLDER)
		g.SetViewOnTop(views.EDIT_FOLDER)
		return nil
	}
}
