package filetree

import (
	"fmt"
	"lazycurl/pkg/collection"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
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

func FileTree(g *gocui.Gui, maxX, maxY int, tree []collection.FileNode, fullScreen bool, addFolderFunc func(foldersPath string) error) (bool, error) {
	if rootNodes == nil {
		rootNodes = tree
	}

	isMenuOpen, err := initFileTreeModal(g, maxX, maxY, fullScreen)
	if err != nil {
		return isMenuOpen, err
	}

	addFolder(g, maxX, maxY, addFolderFunc)

	if err := g.SetKeybinding("", gocui.KeyCtrlSlash, gocui.ModNone, toggleFileTree); err != nil {
		return isMenuOpen, err
	}

	return isMenuOpen, nil
}

func initFileTreeModal(g *gocui.Gui, maxX, maxY int, fullScreen bool) (bool, error) {
	var x0, y0, x1, y1 int

	if fullScreen {
		x0 = 0
		y0 = 0
		x1 = maxX
		y1 = maxY
	} else {
		height := maxY - views.BOTTOM_MESSAGE
		x0 = views.FULL
		x1 = maxX/6 - 2
		y0 = views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
		y1 = height - views.LOGS_BOTTOM
	}

	v, err := g.SetView(views.FILE_TREE_VIEW, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao iniciar file tree:", err)
			return true, err
		}

		v.Title = " [Tab] Files "
		v.Highlight = false
		v.Visible = true
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		rebuildFlatList()
		renderTree(v)

		// Garante que a view sobreponha as outras (Z-index)
		if fullScreen {
			g.SetViewOnTop(views.FILE_TREE_VIEW)
		}

		if err := setupModalKeys(g); err != nil {
			return v.Visible, err
		}
	}

	if utils.ViewHasFocus(g, views.FILE_TREE_VIEW) {
		v.Highlight = true
	} else {
		v.Highlight = false
	}

	return v.Visible, nil
}

func toggleFileTree(g *gocui.Gui, v *gocui.View) error {
	fileView, err := g.View(views.FILE_TREE_VIEW)
	if err != nil {
		return err
	}

	if !fileView.Visible {
		fileView.Visible = true
		g.SetCurrentView(views.FILE_TREE_VIEW)
		g.SetViewOnTop(views.FILE_TREE_VIEW)

		rebuildFlatList()
		renderTree(fileView)
	} else {
		fileView.Visible = false
		g.SetCurrentView(views.METHOD)
	}
	return nil
}

func renderTree(v *gocui.View) {
	oldX, oldY := v.Cursor()
	v.Clear()
	for _, item := range flatItems {
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
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeyEnter, gocui.ModNone, toggleFolder)
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeySpace, gocui.ModNone, toggleFolder)
	g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeyCtrlF, gocui.ModNone, helper.ToggleView(views.ADD_FOLDER))

	return nil
}
