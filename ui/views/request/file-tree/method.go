package filetree

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

var addFileMethods = []string{request.GET, request.POST, request.PUT, request.DELETE, request.HEAD, request.PATCH, request.OPTIONS}

func methodSelectInput(g *gocui.Gui, x0, y0 int) error {
	mv, err := g.SetView(views.ADD_FILE_METHOD, x0+2, y0+2, x0+11, y0+4, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		mv.Title = "Method"
		mv.Frame = true
		mv.FrameColor = gocui.ColorWhite
		mv.TitleColor = gocui.ColorWhite
		mv.Visible = false
		fmt.Fprint(mv, "GET")

		g.SetKeybinding(views.ADD_FILE_METHOD, gocui.KeySpace, gocui.ModNone, openAddFileMethodModal)
		g.SetKeybinding(views.ADD_FILE_METHOD, gocui.KeyEnter, gocui.ModNone, openAddFileMethodModal)
	}

	return nil
}

func openAddFileMethodModal(g *gocui.Gui, v *gocui.View) error {
	methodV, err := g.View(views.ADD_FILE_METHOD)
	if err != nil {
		return err
	}

	x0, _, x1, y1 := methodV.Dimensions()
	modalHeight := len(addFileMethods) + 2
	modalY0 := y1
	modalY1 := modalY0 + modalHeight

	if mv, err := g.SetView("add_file_method_modal", x0, modalY0, x1, modalY1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		mv.FrameColor = gocui.ColorGreen
		mv.TitleColor = gocui.ColorGreen
		mv.Highlight = true
		mv.SelFgColor = gocui.ColorBlack
		mv.SelBgColor = gocui.ColorGreen

		for _, method := range addFileMethods {
			fmt.Fprintln(mv, utils.FormatLineFullWidth(mv, method))
		}

		mv.ViewLinesHeight()

		g.SetKeybinding("add_file_method_modal", gocui.KeyEnter, gocui.ModNone, selectAddFileMethod)
		g.SetKeybinding("add_file_method_modal", gocui.KeyEsc, gocui.ModNone, closeAddFileMethodModal)

		g.SetCurrentView("add_file_method_modal")
	}
	return nil
}

func selectAddFileMethod(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()

	line, err := v.Line(cy)
	if err != nil {
		return err
	}

	g.Update(func(g *gocui.Gui) error {
		methodView, err := g.View(views.ADD_FILE_METHOD)
		if err != nil {
			return err
		}

		methodView.Clear()
		methodView.SetCursor(0, 0)
		methodView.SetOrigin(0, 0)
		fmt.Fprint(methodView, line)

		closeAddFileMethodModal(g, methodView)
		return nil
	})

	return nil
}

func closeAddFileMethodModal(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView("add_file_method_modal")
	g.SetCurrentView(views.ADD_FILE_METHOD)
	return nil
}
