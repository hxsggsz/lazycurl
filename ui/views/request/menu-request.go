package request

import (
	"fmt"
	"lazycurl/pkg/collection"
	"lazycurl/ui/views"
	"strings"

	"github.com/awesome-gocui/gocui"
)


type flatItem struct {
	Node  *collection.FileNode
	Level int
}

var (
	rootNodes []collection.FileNode
	flatItems []flatItem
)

const (
	FileTreeView  = "file_tree_modal"
	FileTreeLabel = "file_tree_label"
)

func FileTree(g *gocui.Gui, maxX, maxY int, tree []collection.FileNode) error {
	if rootNodes == nil {
		rootNodes = tree
	}

	if err := initFileTreeModal(g, maxX, maxY); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, toggleFileTree); err != nil {
		return err
	}

	return nil
}

func initFileTreeModal(g *gocui.Gui, maxX, maxY int) error {
	w, h := int(float64(maxX)*0.6), int(float64(maxY)*0.6)
	x0 := (maxX - w) / 2
	y0 := (maxY - h) / 2
	x1, y1 := x0+w, y0+h

	if v, err := g.SetView(FileTreeView, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = " [2] Files "
		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.Visible = false

		if err := setupModalKeys(g); err != nil {
			return err
		}
	}
	return nil
}

func toggleFileTree(g *gocui.Gui, v *gocui.View) error {
	fileView, err := g.View(FileTreeView)
	if err != nil {
		return err
	}

	if !fileView.Visible {
		fileView.Visible = true
		g.SetCurrentView(FileTreeView)
		g.SetViewOnTop(FileTreeView)
		rebuildFlatList()
		renderTree(fileView)
	} else {
		fileView.Visible = false
		g.SetCurrentView("main")
	}
	return nil
}

func renderTree(v *gocui.View) {
	oldX, oldY := v.Cursor()
	v.Clear()
	for _, item := range flatItems {
		indent := strings.Repeat("  ", item.Level)
		if !item.Node.IsDir {
			fmt.Fprintf(v, "%s  %s\n", indent, item.Node.Name)
			continue
		}

		icon := "▶"
		if item.Node.Open {
			icon = "▼"
		}

		coloredDir := views.GREEN + indent + icon + " " + item.Node.Name + views.RESET
		fmt.Fprintf(v, "%s\n", coloredDir)
	}

	v.SetCursor(oldX, oldY)
}

func rebuildFlatList() {
	flatItems = []flatItem{}

	var flatten func(nodes []collection.FileNode, level int)
	flatten = func(nodes []collection.FileNode, level int) {
		for i := range nodes {
			flatItems = append(flatItems, flatItem{Node: &nodes[i], Level: level})
			if nodes[i].IsDir && nodes[i].Open {
				flatten(nodes[i].Children, level+1)
			}
		}
	}

	flatten(rootNodes, 0)
}

func toggleFolder(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	_, oy := v.Origin()
	idx := cy + oy

	if idx >= 0 && idx < len(flatItems) {
		if flatItems[idx].Node.IsDir {
			flatItems[idx].Node.Open = !flatItems[idx].Node.Open
			rebuildFlatList()
			renderTree(v)
		}
	}

	return nil
}

func setupModalKeys(g *gocui.Gui) error {
	g.SetKeybinding(FileTreeView, gocui.KeyEnter, gocui.ModNone, toggleFolder)
	g.SetKeybinding(FileTreeView, gocui.KeySpace, gocui.ModNone, toggleFolder)

	return nil
}
