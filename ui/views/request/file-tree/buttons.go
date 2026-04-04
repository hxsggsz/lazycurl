package filetree

import (
	"fmt"
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func addFileButtons(g *gocui.Gui, x0, x1, y0 int) error {
	btnV, err := g.SetView(views.ADD_FILE_BUTTONS, x0, y0+8, x1, y0+10, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		btnV.Visible = false
		btnV.Frame = false
	}
	return nil
}

func renderAddFileButtons(g *gocui.Gui) {
	v, err := g.View(views.ADD_FILE_BUTTONS)
	if err != nil {
		return
	}

	cancelStyle := "  Cancel"
	createStyle := "  Create"

	if addFileFocusedIdx == addFileFocusCancel {
		cancelStyle = views.RED + " [Cancel]" + views.RESET
	}

	if addFileFocusedIdx == addFileFocusCreate {
		createStyle = views.GREEN + " [Create]" + views.RESET
	}

	v.Clear()
	fmt.Fprintf(v, "%s    %s", cancelStyle, createStyle)
}

func handleAddFileButtonClick(fm *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		cx, _ := v.Cursor()
		if cx < 10 {
			addFileFocusedIdx = addFileFocusCancel
		} else {
			addFileFocusedIdx = addFileFocusCreate
		}
		renderAddFileButtons(g)
		return confirmAddFile(fm)(g, v)
	}
}
