package filemanager

import (
	"fmt"
	"lazycurl/pkg/collection"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"strings"

	"github.com/awesome-gocui/gocui"
)

// FileManager manages the file tree state and provides methods for rendering
// and updating the tree view. It bridges the collection data with the UI.
type FileManager struct {
	// Collection is the source of truth for the file tree data
	Collection *collection.Collection
	// flatItems holds the flattened tree for linear rendering
	flatItems []flatItem
}

type flatItem struct {
	Node  *collection.FileNode
	Level int
}


func (fm *FileManager) ToggleFileTree(g *gocui.Gui, v *gocui.View) error {
	fileView, err := g.View(views.FILE_TREE_VIEW)
	if err != nil {
		return err
	}

	if !fileView.Visible {
		fileView.Visible = true
		g.SetCurrentView(views.FILE_TREE_VIEW)
		g.SetViewOnTop(views.FILE_TREE_VIEW)

		fm.UpdateTree(g)
	} else {
		fileView.Visible = false
		g.SetCurrentView(views.METHOD)
	}
	return nil
}

// renderTree clears and re-renders the file tree view from the flat item list.
func (fm *FileManager) RenderTree(v *gocui.View) {
	oldX, oldY := v.Cursor()
	v.Clear()
	for _, item := range fm.flatItems {
		indent := strings.Repeat("  ", item.Level)
		if !item.Node.IsDir {

			fileLine := indent + " " + item.Node.Name
			formatedFile := utils.FormatLineFullWidth(v, fileLine)
			fmt.Fprintf(v, "%s\n", formatedFile)
			continue
		}

		icon := "▶"
		if item.Node.Open {
			icon = "▼"
		}

		coloredDir := views.GREEN + indent + " " + icon + " " + item.Node.Name + views.RESET
		coloredDir = utils.FormatLineFullWidth(v, coloredDir)
		fmt.Fprintf(v, "%s\n", coloredDir)
	}

	v.SetCursor(oldX, oldY)
}

// rebuildFlatList flattens the hierarchical file tree into a linear slice
// for rendering, respecting the Open state of directory nodes.
func (fm *FileManager) RebuildFlatList() {
	fm.flatItems = []flatItem{}

	var flatten func(nodes []collection.FileNode, level int)
	flatten = func(nodes []collection.FileNode, level int) {
		for i := range nodes {
			fm.flatItems = append(fm.flatItems, flatItem{Node: &nodes[i], Level: level})
			if nodes[i].IsDir && nodes[i].Open {
				flatten(nodes[i].Children, level+1)
			}
		}
	}

	flatten(fm.Collection.Files, 0)
}

// toggleFolder expands or collapses a directory node at the current cursor position.
func (fm *FileManager) ToggleFolder(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	_, oy := v.Origin()
	idx := cy + oy

	if idx >= 0 && idx < len(fm.flatItems) {
		if fm.flatItems[idx].Node.IsDir {
			fm.flatItems[idx].Node.Open = !fm.flatItems[idx].Node.Open
			fm.RebuildFlatList()
			fm.RenderTree(v)
		}
	}

	return nil
}

// setupModalKeys registers keybindings for the file tree view.
func (fm *FileManager) SetupModalKeys(g *gocui.Gui) error {
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeyEnter, gocui.ModNone, fm.ToggleFolder)
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeySpace, gocui.ModNone, fm.ToggleFolder)
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeyCtrlF, gocui.ModNone, helper.ToggleView(views.ADD_FOLDER))

	return nil
}

// UpdateTree reloads the collection from disk, preserves open folder state,
// rebuilds the flat list, and re-renders the tree view.
func (fm *FileManager) UpdateTree(g *gocui.Gui) {
	openPaths := fm.Collection.GetOpenPaths()
	fm.Collection.LoadCollectionFiles()
	fm.Collection.RestoreOpenPaths(openPaths)
	fm.RebuildFlatList()
	if v, err := g.View(views.FILE_TREE_VIEW); err == nil {
		fm.RenderTree(v)
	}
}

