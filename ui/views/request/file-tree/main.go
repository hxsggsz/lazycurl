package filetree

import (
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

// FileTree initializes and manages the file tree view in the UI.
// It registers keybindings for navigation and folder creation.
func FileTree(g *gocui.Gui, maxX, maxY int, fm *fm.FileManager, fullScreen bool) (bool, error) {
	isMenuOpen, err := initFileTreeModal(g, maxX, maxY, fullScreen, fm)
	if err != nil {
		return isMenuOpen, err
	}

	addFolder(g, maxX, maxY, fm)
	deleteNode(g, maxX, maxY, fm)

	if err := g.SetKeybinding("", gocui.KeyCtrlSlash, gocui.ModNone, fm.ToggleFileTree); err != nil {
		return isMenuOpen, err
	}

	return isMenuOpen, nil
}

func initFileTreeModal(g *gocui.Gui, maxX, maxY int, fullScreen bool, fm *fm.FileManager) (bool, error) {
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

		fm.RebuildFlatList()
		fm.RenderTree(v)

		if fullScreen {
			g.SetViewOnTop(views.FILE_TREE_VIEW)
		}

		if err := fm.SetupModalKeys(g); err != nil {
			return v.Visible, err
		}

		g.SetKeybinding(views.FILE_TREE_VIEW, gocui.KeyDelete, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			renderDeleteConfirm(g, fm)
			return nil
		})
	}

	if utils.ViewHasFocus(g, views.FILE_TREE_VIEW) {
		v.Highlight = true
	} else {
		v.Highlight = false
	}

	return v.Visible, nil
}
