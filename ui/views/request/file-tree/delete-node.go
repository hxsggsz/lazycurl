package filetree

import (
	"fmt"
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
	"time"

	"github.com/awesome-gocui/gocui"
)

var deleteConfirmSelected = 0

// deleteNode creates the confirmation modal view for deleting a file or folder.
func deleteNode(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	width := 40
	x0 := (maxX - width) / 2
	x1 := x0 + width

	y0 := (maxY - 6) / 2
	y1 := y0 + 6

	v, err := g.SetView(views.DELETE_NODE, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Confirm Delete"
		v.Frame = true
		v.Wrap = false
		v.Visible = false

		v.FrameColor = gocui.ColorRed
		v.TitleColor = gocui.ColorRed
	}

	g.SetKeybinding(views.DELETE_NODE, gocui.KeyArrowLeft, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		deleteConfirmSelected = 0
		renderDeleteConfirm(g, fileManager)
		return nil
	})
	g.SetKeybinding(views.DELETE_NODE, gocui.KeyArrowRight, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		deleteConfirmSelected = 1
		renderDeleteConfirm(g, fileManager)
		return nil
	})
	g.SetKeybinding(views.DELETE_NODE, gocui.KeyTab, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		deleteConfirmSelected = (deleteConfirmSelected + 1) % 2
		renderDeleteConfirm(g, fileManager)
		return nil
	})

	g.SetKeybinding(views.DELETE_NODE, gocui.KeyEnter, gocui.ModNone, confirmDelete(fileManager))
	g.SetKeybinding(views.DELETE_NODE, gocui.KeyCtrlQ, gocui.ModNone, helper.CloseView(views.DELETE_NODE))
	g.SetKeybinding(views.DELETE_NODE, gocui.KeyEsc, gocui.ModNone, helper.CloseView(views.DELETE_NODE))
	return nil
}

func renderDeleteConfirm(g *gocui.Gui, fileManager *fm.FileManager) {
	node, err := fileManager.GetSelectedNode(g)
	if err != nil {
		return
	}

	v, err := g.View(views.DELETE_NODE)
	if err != nil {
		return
	}

	v.Clear()
	fmt.Fprintln(v)
	fmt.Fprintf(v, "  Delete %q?\n\n", node.Name)

	yesStyle := "  Yes"
	noStyle := "  No"

	if deleteConfirmSelected == 0 {
		yesStyle = views.RED + " [Yes]" + views.RESET
	} else {
		yesStyle = "  Yes"
	}

	if deleteConfirmSelected == 1 {
		noStyle = views.GREEN + " [No]" + views.RESET
	} else {
		noStyle = "  No"
	}

	fmt.Fprintf(v, "%s    %s\n", yesStyle, noStyle)

	v.Visible = true
	g.SetViewOnTop(views.DELETE_NODE)
	g.SetCurrentView(views.DELETE_NODE)
}

func confirmDelete(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if deleteConfirmSelected == 1 {
			helper.CloseView(views.DELETE_NODE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			return nil
		}

		if err := fileManager.Collection.DeletePath(node.Name); err != nil {
			log.Printf("failed to delete %s: %v", node.Name, err)
			views.ShowToast(g, fmt.Sprintf("failed to delete: %v", err), "error", 2*time.Second)
			helper.CloseView(views.DELETE_NODE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		helper.CloseView(views.DELETE_NODE)(g, v)
		views.ShowToast(g, fmt.Sprintf("%q deleted", node.Name), "success", 2*time.Second)
		fileManager.UpdateTree(g)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		return nil
	}
}
