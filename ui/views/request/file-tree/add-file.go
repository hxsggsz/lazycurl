package filetree

import (
	fm "lazycurl/pkg/file-manager"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var addFileSubViews = []string{
	views.ADD_FILE_URL,
	views.ADD_FILE_METHOD,
	views.ADD_FILE_NAME,
}

func addFile(g *gocui.Gui, maxX, maxY int, fileManager *fm.FileManager) error {
	width := 40
	height := 9
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
	requestNameInput(g, x0, x1, y0)
	methodSelectInput(g, x0, y0)
	requestUrlInput(g, x0, x1, y0)

	for _, name := range addFileSubViews {
		g.SetKeybinding(name, gocui.KeyArrowRight, gocui.ModNone, nextAddFileInput())
		g.SetKeybinding(name, gocui.KeyArrowLeft, gocui.ModNone, prevAddFileInput())
		g.SetKeybinding(name, gocui.KeyTab, gocui.ModNone, nextAddFileInput())
		g.SetKeybinding(name, gocui.KeyTab, gocui.ModShift, prevAddFileInput())
	}

	g.SetKeybinding(views.ADD_FILE, gocui.KeyCtrlQ, gocui.ModNone, closeAddFileModal)
	g.SetKeybinding(views.ADD_FILE, gocui.KeyEsc, gocui.ModNone, closeAddFileModal)

	return nil
}

func focusAddFileInput(idx int) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		for _, name := range addFileSubViews {
			if sv, err := g.View(name); err == nil {
				sv.FrameColor = gocui.ColorWhite
				sv.TitleColor = gocui.ColorWhite
			}
		}

		target := addFileSubViews[idx]
		if tv, err := g.View(target); err == nil {
			tv.FrameColor = gocui.ColorGreen
			tv.TitleColor = gocui.ColorGreen
			g.SetCurrentView(target)
		}
		return nil
	}
}

func focusAddFileInputByName(name string) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		for _, n := range addFileSubViews {
			if sv, err := g.View(n); err == nil {
				sv.FrameColor = gocui.ColorWhite
				sv.TitleColor = gocui.ColorWhite
			}
		}

		if tv, err := g.View(name); err == nil {
			tv.FrameColor = gocui.ColorGreen
			tv.TitleColor = gocui.ColorGreen
			g.SetCurrentView(name)
		}
		return nil
	}
}

func nextAddFileInput() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		for i, name := range addFileSubViews {
			if name == v.Name() {
				next := (i + 1) % len(addFileSubViews)
				return focusAddFileInput(next)(g, v)
			}
		}
		return nil
	}
}

func prevAddFileInput() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		for i, name := range addFileSubViews {
			if name == v.Name() {
				prev := (i - 1 + len(addFileSubViews)) % len(addFileSubViews)
				return focusAddFileInput(prev)(g, v)
			}
		}
		return nil
	}
}

func setAddFileSubViewsVisible(g *gocui.Gui, visible bool) error {
	for _, name := range addFileSubViews {
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
			g.DeleteView("add_file_method_modal")
			g.SetCurrentView(views.FILE_TREE_VIEW)
			return nil
		}

		outerV.Visible = true
		setAddFileSubViewsVisible(g, true)

		g.SetViewOnTop(views.ADD_FILE)
		for _, name := range addFileSubViews {
			g.SetViewOnTop(name)
		}

		focusAddFileInput(0)(g, v)

		return nil
	}
}

func closeAddFileModal(g *gocui.Gui, v *gocui.View) error {
	if outerV, err := g.View(views.ADD_FILE); err == nil {
		outerV.Visible = false
	}
	setAddFileSubViewsVisible(g, false)
	g.DeleteView("add_file_method_modal")
	g.SetCurrentView(views.FILE_TREE_VIEW)
	return nil
}
