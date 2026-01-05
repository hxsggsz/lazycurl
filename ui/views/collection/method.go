package collection

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var methods = []string{request.GET, request.POST, request.PUT, request.DELETE, request.HEAD, request.PATCH, request.OPTIONS}

func Method(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.METHOD

	x0 := views.FULL
	x1 := 10 - views.LAYOUT_SECTION_Y_GAP

	y0 := 0
	y1 := views.LAYOUT_INPUT_HEIGHT

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf("[%d]", 1)
		v.FrameColor = gocui.ColorWhite
		fmt.Fprint(v, "GET")

		if err := g.SetKeybinding(viewName, gocui.KeySpace, gocui.ModNone, openMethodModal); err != nil {
			return err
		}

		if err := g.SetKeybinding(viewName, gocui.KeyEnter, gocui.ModNone, openMethodModal); err != nil {
			return err
		}

		g.SetViewOnBottom(viewName)
	}

	return nil
}

func GetCurrentMethod(g *gocui.Gui) string {
	v, err := g.View(views.METHOD)
	if err != nil {
		return "GET"
	}
	return strings.TrimSpace(v.Buffer())
}

func openMethodModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	// Modal centralizado
	modalWidth := 20
	modalHeight := 8
	x0 := (maxX - modalWidth) / 2
	y0 := (maxY - modalHeight) / 2
	x1 := x0 + modalWidth
	y1 := y0 + modalHeight

	if v, err := g.SetView("method_modal", x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		g.Cursor = false

		v.Title = "Select Method"
		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen

		v.Highlight = true
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen

		for _, method := range methods {
			fmt.Fprintln(v, method)
		}

		v.ViewLinesHeight()

		if err := registerMethodViewNavigation(g); err != nil {
			return err
		}

		g.SetCurrentView("method_modal")

	}
	return nil
}

func registerMethodViewNavigation(g *gocui.Gui) error {
	methodModalKeymaps := utils.KeybindsMaps{
		{Key: gocui.KeyEsc, Modifier: gocui.ModNone}:   closeMethodModal,
		{Key: gocui.KeyEnter, Modifier: gocui.ModNone}: selectMethod,
	}

	if err := utils.SetKeybind(g, methodModalKeymaps, "method_modal"); err != nil {
		return err
	}

	return nil
}

func selectMethod(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		return err
	}

	log.Printf("Selected method: %s", line)

	g.Update(func(g *gocui.Gui) error {
		methodView, err := g.View(views.METHOD)
		if err != nil {
			return err
		}

		methodView.Clear()
		methodView.SetCursor(0, 0)
		methodView.SetOrigin(0, 0)
		fmt.Fprint(methodView, line)

		closeMethodModal(g, methodView)
		return nil
	})

	return nil
}

func closeMethodModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView("method_modal")
	g.SetCurrentView(views.METHOD)
	return nil
}

func moveDownModal(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		if v == nil {
			return nil
		}

		cx, cy := v.Cursor()

		if cy < len(methods)-1 {
			if err := v.SetCursor(cx, cy+1); err != nil {
				return err
			}
		}

		cx, cy = v.Cursor()
		return nil
	})

	return nil
}

func moveUpModal(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		if v == nil {
			return nil
		}

		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil {
			ox, oy := v.Origin()
			v.SetOrigin(ox, oy-1)
		}

		return nil
	})
	return nil
}
