package collection

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

func Method(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.METHOD

	x0 := 0
	x1 := maxX / 16

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

func openMethodModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()

	// Modal centralizado
	modalWidth := 20
	modalHeight := 9
	x0 := (maxX - modalWidth) / 2
	y0 := (maxY - modalHeight) / 2
	x1 := x0 + modalWidth
	y1 := y0 + modalHeight

	if v, err := g.SetView("methodModal", x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Select Method"
		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true

		methods := []string{request.GET, request.POST, request.PUT, request.DELETE, request.HEAD, request.PATCH, request.OPTIONS}
		for _, method := range methods {
			fmt.Fprintln(v, method)
		}

		if err := RegisterGlobalViewNavigation(g); err != nil {
			return err
		}
		// Enter: seleciona e fecha
		if err := g.SetKeybinding("methodModal", gocui.KeyEnter, gocui.ModNone, selectMethod); err != nil {
			return err
		}

		// ESC: fecha sem selecionar
		if err := g.SetKeybinding("methodModal", gocui.KeyEsc, gocui.ModNone, closeMethodModal); err != nil {
			return err
		}

		if err := g.SetKeybinding("methodModal", gocui.KeyEsc, gocui.ModNone, closeMethodModal); err != nil {
			return err
		}

		// Navegação
		if err := g.SetKeybinding("methodModal", gocui.KeyArrowDown, gocui.ModNone, moveDownModal); err != nil {
			return err
		}
		if err := g.SetKeybinding("methodModal", gocui.KeyArrowUp, gocui.ModNone, moveUpModal); err != nil {
			return err
		}

		g.SetCurrentView("methodModal")
	}
	return nil
}

func RegisterGlobalViewNavigation(g *gocui.Gui) error {
	methodModalKeymaps := map[gocui.Key]func(g *gocui.Gui, v *gocui.View) error{
		'q': closeMethodModal, gocui.KeyEsc: closeMethodModal,
		gocui.KeyEnter: selectMethod, gocui.KeyArrowDown: moveDownModal,
		gocui.KeyArrowUp: moveUpModal,
	}

	for key, handler := range methodModalKeymaps {
		if err := g.SetKeybinding(views.METHOD, key, gocui.ModNone, handler); err != nil {
			return err
		}
	}

	return nil
}

func selectMethod(g *gocui.Gui, v *gocui.View) error {
	// Pega a posição do cursor
	_, cy := v.Cursor()

	// Pega a linha na posição do cursor
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
	g.DeleteView("methodModal")
	g.SetCurrentView(views.METHOD)
	return nil
}

func moveDownModal(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		v.SetOrigin(ox, oy+1)
	}
	return nil
}

func moveUpModal(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	if cy > 0 {
		v.SetCursor(cx, cy-1)
	} else {
		ox, oy := v.Origin()
		if oy > 0 {
			v.SetOrigin(ox, oy-1)
		}
	}
	return nil
}
