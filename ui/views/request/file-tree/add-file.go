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

var addFileInputViews = []string{
	views.ADD_FILE_METHOD,
	views.ADD_FILE_NAME,
	views.ADD_FILE_URL,
}

const (
	addFileFocusMethod = iota
	addFileFocusName
	addFileFocusURL
	addFileFocusCancel
	addFileFocusCreate
)

var addFileFocusCount = 5
var addFileFocusedIdx = addFileFocusMethod

func addFile(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	width := 40
	height := 11
	x0 := (maxX - width) / 2
	y0 := (maxY - height) / 2
	x1 := x0 + width
	y1 := y0 + height

	v, err := g.SetView(views.ADD_FILE, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "New Request File"
		v.Frame = true
		v.FrameColor = gocui.ColorCyan
		v.TitleColor = gocui.ColorCyan
		v.Visible = false
		g.SetViewOnBottom(views.ADD_FILE)
	}

	requestUrlInput(g, x0, x1, y0)
	methodSelectInput(g, x0, y0)
	requestNameInput(g, x0, x1, y0)
	addFileButtons(g, x0, x1, y0)

	for _, name := range addFileInputViews {
		g.SetKeybinding(name, gocui.KeyArrowRight, gocui.ModNone, nextAddFileFocus())
		g.SetKeybinding(name, gocui.KeyArrowLeft, gocui.ModNone, prevAddFileFocus())
		g.SetKeybinding(name, gocui.KeyTab, gocui.ModNone, nextAddFileFocus())
		g.SetKeybinding(name, gocui.KeyTab, gocui.ModShift, prevAddFileFocus())
	}

	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.KeyArrowRight, gocui.ModNone, nextAddFileFocus())
	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.KeyArrowLeft, gocui.ModNone, prevAddFileFocus())
	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.KeyTab, gocui.ModNone, nextAddFileFocus())
	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.KeyTab, gocui.ModShift, prevAddFileFocus())
	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.MouseLeft, gocui.ModNone, handleAddFileButtonClick(fileManager))

	g.SetKeybinding(views.ADD_FILE, gocui.KeyArrowRight, gocui.ModNone, nextAddFileFocus())
	g.SetKeybinding(views.ADD_FILE, gocui.KeyArrowLeft, gocui.ModNone, prevAddFileFocus())
	g.SetKeybinding(views.ADD_FILE, gocui.KeyTab, gocui.ModNone, nextAddFileFocus())
	g.SetKeybinding(views.ADD_FILE, gocui.KeyTab, gocui.ModShift, prevAddFileFocus())

	g.SetKeybinding(views.ADD_FILE, gocui.KeyEnter, gocui.ModNone, confirmAddFile(fileManager))
	g.SetKeybinding(views.ADD_FILE_BUTTONS, gocui.KeyEnter, gocui.ModNone, confirmAddFile(fileManager))
	g.SetKeybinding(views.ADD_FILE, gocui.KeyCtrlQ, gocui.ModNone, closeAddFileModal)
	g.SetKeybinding(views.ADD_FILE, gocui.KeyEsc, gocui.ModNone, closeAddFileModal)

	return nil
}

func setAddFileFocus(idx int, g *gocui.Gui) error {
	addFileFocusedIdx = idx

	for _, name := range addFileInputViews {
		if sv, err := g.View(name); err == nil {
			sv.FrameColor = gocui.ColorWhite
			sv.TitleColor = gocui.ColorWhite
		}
	}

	if idx < len(addFileInputViews) {
		target := addFileInputViews[idx]
		if tv, err := g.View(target); err == nil {
			tv.FrameColor = gocui.ColorGreen
			tv.TitleColor = gocui.ColorGreen
			g.SetCurrentView(target)
		}
	} else {
		g.SetCurrentView(views.ADD_FILE_BUTTONS)
	}

	renderAddFileButtons(g)
	return nil
}

func nextAddFileFocus() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		next := (addFileFocusedIdx + 1) % addFileFocusCount
		return setAddFileFocus(next, g)
	}
}

func prevAddFileFocus() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		prev := (addFileFocusedIdx - 1 + addFileFocusCount) % addFileFocusCount
		return setAddFileFocus(prev, g)
	}
}

func setAddFileSubViewsVisible(g *gocui.Gui, visible bool) error {
	for _, name := range addFileInputViews {
		v, err := g.View(name)
		if err != nil {
			continue
		}
		v.Visible = visible
	}
	return nil
}

func toggleAddFileModal() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		outerV, err := g.View(views.ADD_FILE)
		if err != nil {
			return err
		}

		if outerV.Visible {
			outerV.Visible = false
			setAddFileSubViewsVisible(g, false)
			if btnV, err := g.View(views.ADD_FILE_BUTTONS); err == nil {
				btnV.Visible = false
			}
			g.DeleteView("add_file_method_modal")
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		outerV.Visible = true
		setAddFileSubViewsVisible(g, true)

		if btnV, err := g.View(views.ADD_FILE_BUTTONS); err == nil {
			btnV.Visible = true
		}

		g.SetViewOnTop(views.ADD_FILE)
		for _, name := range addFileInputViews {
			g.SetViewOnTop(name)
		}
		g.SetViewOnTop(views.ADD_FILE_BUTTONS)

		addFileFocusedIdx = addFileFocusMethod
		setAddFileFocus(addFileFocusMethod, g)

		return nil
	}
}

func closeAddFileModal(g *gocui.Gui, v *gocui.View) error {
	if outerV, err := g.View(views.ADD_FILE); err == nil {
		outerV.Visible = false
	}
	setAddFileSubViewsVisible(g, false)
	if btnV, err := g.View(views.ADD_FILE_BUTTONS); err == nil {
		btnV.Visible = false
	}
	g.DeleteView("add_file_method_modal")
	g.SetCurrentView(views.FILE_TREE_VIEW)
	return nil
}

func confirmAddFile(fileManager *fm.FileManager) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if addFileFocusedIdx == addFileFocusCancel {
			closeAddFileModal(g, v)
			return nil
		}

		urlV, _ := g.View(views.ADD_FILE_URL)
		methodV, _ := g.View(views.ADD_FILE_METHOD)
		fileNameV, _ := g.View(views.ADD_FILE_NAME)

		if urlV == nil || methodV == nil || fileNameV == nil {
			return nil
		}

		url := urlV.Buffer()
		method := methodV.Buffer()
		fileName := fileNameV.Buffer()

		everyInputsAreFilled := url != "" || method != "" || fileName != ""

		if addFileFocusedIdx != addFileFocusCreate || !everyInputsAreFilled {
			return nil
		}

		if fileName == "" {
			views.ShowToast(g, "Request name is required", "error", 2*time.Second)
			return nil
		}

		node, err := fileManager.GetSelectedNode(g)
		if err != nil {
			views.ShowToast(g, "no item selected", "error", 2*time.Second)
			helper.CloseView(views.ADD_FILE)(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		if err := fileManager.Collection.AddFileWithNode(fileName, method, url, node.Path); err != nil {
			log.Printf("failed to create request file: %v", err)
			views.ShowToast(g, fmt.Sprintf("failed to create: %v", err), "error", 2*time.Second)
			closeAddFileModal(g, v)
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		closeAddFileModal(g, v)
		views.ShowToast(g, fmt.Sprintf("%q created", fileName), "success", 2*time.Second)
		fileManager.UpdateTree(g)
		g.SetCurrentView(views.FILE_TREE_VIEW)
		return nil
	}
}
