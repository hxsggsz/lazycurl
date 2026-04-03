package filetree

import (
	"fmt"
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/awesome-gocui/gocui"
)

// editFile creates the modal view for renaming a file.
func editFile(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView(views.EDIT_FILE, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		v.Title = " Rename File "
		v.Frame = true
		v.Wrap = true
		v.Editable = true
		v.Visible = false

		v.FrameColor = gocui.ColorYellow
		v.TitleColor = gocui.ColorYellow
	}

	g.SetKeybinding(views.EDIT_FILE, gocui.KeyEnter, gocui.ModNone, renameFile(fileManager))
	g.SetKeybinding(views.EDIT_FILE, gocui.KeyCtrlQ, gocui.ModNone, helper.CloseView(views.EDIT_FILE))
	g.SetKeybinding(views.EDIT_FILE, gocui.KeyEsc, gocui.ModNone, helper.CloseView(views.EDIT_FILE))
	return nil
}

// getEditFileInput retrieves the current text from the edit file modal.
func getEditFileInput(g *gocui.Gui) string {
	v, err := g.View(views.EDIT_FILE)
	if err != nil {
		log.Println("Erro ao obter view do modal:", err)
		return ""
	}
	return v.Buffer()
}

// clearEditFileInput clears the text in the edit file modal.
func clearEditFileInput(g *gocui.Gui) {
	if v, err := g.View(views.EDIT_FILE); err == nil {
		v.Clear()
	}
}

// populateEditFileInput sets the initial file name in the edit modal.
func populateEditFileInput(g *gocui.Gui, name string) {
	if v, err := g.View(views.EDIT_FILE); err == nil {
		v.Clear()
		fmt.Fprint(v, name)
	}
}

// renameFile returns a keybinding handler that renames the selected file.
func renameFile(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		newName := getEditFileInput(g)
		if newName == "" {
			log.Println("file name cannot be empty")
			return nil
		}

		if !strings.HasSuffix(newName, ".json") {
			newName = newName + ".json"
		}

		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		if node.IsDir {
			views.ShowToast(g, "selected item is not a file", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		rootPath := fileManager.Collection.GetRootPath()
		relPath, err := filepath.Rel(rootPath, node.Path)
		if err != nil {
			log.Printf("failed to get relative path: %v", err)
			views.ShowToast(g, "failed to rename file", "error", 2*time.Second)
			helper.CloseView(views.EDIT_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		if err := fileManager.Collection.RenameNode(relPath, newName); err != nil {
			log.Printf("failed to rename file: %v", err)
			views.ShowToast(g, fmt.Sprintf("failed to rename: %v", err), "error", 2*time.Second)
			helper.CloseView(views.EDIT_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		clearEditFileInput(g)
		helper.CloseView(views.EDIT_FILE)(g, v)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		views.ShowToast(g, fmt.Sprintf("%q renamed to %q", node.Name, newName), "success", 2*time.Second)
		fileManager.UpdateTree(g)
		return nil
	}
}

// triggerEditFile opens the edit file modal with the current file name.
func triggerEditFile(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			return nil
		}

		if node.IsDir {
			views.ShowToast(g, "selected item is not a file", "error", 2*time.Second)
			return nil
		}

		nameWithoutExt := strings.TrimSuffix(node.Name, ".json")
		populateEditFileInput(g, nameWithoutExt)

		if modalView, err := g.View(views.EDIT_FILE); err == nil {
			modalView.Visible = true
		}

		g.SetCurrentView(views.EDIT_FILE)
		g.SetViewOnTop(views.EDIT_FILE)
		return nil
	}
}
