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

// addFile creates the modal view for entering a new file name.
func addFile(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView(views.ADD_FILE, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		v.Title = " File Name "
		v.Subtitle = " Enter file name (without .json extension) "
		v.Frame = true
		v.Wrap = true
		v.Editable = true
		v.Visible = false

		v.FrameColor = gocui.ColorCyan
		v.TitleColor = gocui.ColorCyan
	}

	g.SetKeybinding(views.ADD_FILE, gocui.KeyEnter, gocui.ModNone, createFile(fileManager))
	g.SetKeybinding(views.ADD_FILE, gocui.KeyCtrlQ, gocui.ModNone, helper.CloseView(views.ADD_FILE))
	g.SetKeybinding(views.ADD_FILE, gocui.KeyEsc, gocui.ModNone, helper.CloseView(views.ADD_FILE))
	return nil
}

// getFileNameInput retrieves the current text from the add file modal.
func getFileNameInput(g *gocui.Gui) string {
	v, err := g.View(views.ADD_FILE)
	if err != nil {
		log.Println("Erro ao obter view do modal:", err)
		return ""
	}
	return v.Buffer()
}

// clearFileNameInput clears the text in the add file modal.
func clearFileNameInput(g *gocui.Gui) {
	if v, err := g.View(views.ADD_FILE); err == nil {
		v.Clear()
	}
}

// createFile returns a keybinding handler that creates the entered file
// with a default JSON request structure, closes the modal, shows a success
// toast, and updates the tree view.
func createFile(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		fileName := getFileNameInput(g)
		if fileName == "" {
			log.Println("file name cannot be empty")
			return nil
		}

		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			helper.CloseView(views.ADD_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		rootPath := fileManager.Collection.GetRootPath()
		var relPath string

		if node.IsDir {
			rel, err := filepath.Rel(rootPath, node.Path)
			if err != nil || rel == "." {
				relPath = fileName
			} else {
				relPath = rel + "/" + fileName
			}
		} else {
			parentDir := filepath.Dir(node.Path)
			rel, err := filepath.Rel(rootPath, parentDir)
			if err != nil || rel == "." {
				relPath = fileName
			} else {
				relPath = rel + "/" + fileName
			}
		}

		if err := fileManager.Collection.AddFile(relPath); err != nil {
			log.Printf("failed to create file: %v", err)
			views.ShowToast(g, fmt.Sprintf("failed to create file: %v", err), "error", 2*time.Second)
			helper.CloseView(views.ADD_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		clearFileNameInput(g)
		helper.CloseView(views.ADD_FILE)(g, v)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		views.ShowToast(g, fmt.Sprintf("%q created", fileName), "success", 2*time.Second)
		fileManager.UpdateTree(g)
		return nil
	}
}
