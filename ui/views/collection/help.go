package collection

import (
	"fmt"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"

	"github.com/awesome-gocui/gocui"
)

const (
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
	BgSelected  = "\033[47;30m" // Fundo branco, texto preto (estilo seleção)
)

func Help(g *gocui.Gui, maxX, maxY int) error {
	viewWidth := 18
	x0, y0 := maxX-viewWidth-1, maxY-2
	x1, y1 := maxX-1, maxY

	if v, err := g.SetView("help_label", x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		ShowHelpModal(g, v)

		v.Frame = false

		fmt.Fprintf(v, "%spress ? for help%s", views.GREEN, views.RESET)

		g.SetViewOnTop("help_label")

		if err := g.SetKeybinding("", '?', gocui.ModNone, helper.ToggleView("help_modal")); err != nil {
			return err
		}

		if err := g.SetKeybinding("help_modal", gocui.KeyEsc, gocui.ModNone, helper.CloseView("help_modal")); err != nil {
			return err
		}
	}

	return nil
}
func ShowHelpModal(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	w, h := 60, 20
	x0, y0 := (maxX-w)/2, (maxY-h)/2
	x1, y1 := x0+w, y0+h

	if v, err := g.SetView("help_modal", x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		g.SetViewOnBottom("help_modal")

		v.Title = " Keybindings "
		v.FrameColor = gocui.ColorCyan
		v.SelFgColor = gocui.ColorBlack
		v.SelBgColor = gocui.ColorGreen
		v.Highlight = true

		printCategory := func(name string, breakLine ...bool) {
			shouldBreak := true
			if len(breakLine) > 0 {
				shouldBreak = breakLine[0]
			}

			if shouldBreak {
				fmt.Fprintln(v)
			}

			fmt.Fprintf(v, "%s  -- %s --%s\n", ColorBlue, name, ColorReset)
		}

		printKeybind := func(key, desc string) {
			fmt.Fprintf(v, "  %s%-10s%s %s\n", ColorGreen, key, ColorReset, desc)
		}

		printCategory("Global", false)
		printKeybind("<Enter>", "Submit Request")
		printKeybind("<F10>", "Toggle Logs")
		printKeybind("<1>", "Focus in select method view")
		printKeybind("<2>", "Focus in URL view")
		printKeybind("<3>", "Focus in body/headers view")
		printKeybind("<4>", "Focus in response/headers view")
		printKeybind("<Esc>", "Unfocus view")
		printKeybind("<C-c>", "Quit")

		printCategory("Modals")
		printKeybind("<Esc>", "Close modal")
		printKeybind("<q>", "Close modal")

		printCategory("Navigation")
		printKeybind("<up>", "Scroll up")
		printKeybind("<down>", "Scroll down")
		printKeybind("<k>", "Scroll up")
		printKeybind("<j>", "Scroll down")
		printKeybind("<S-ArrowRight>", "Go to next tab")
		printKeybind("<S-ArrowLeft>", "Go to previous tab")

		printCategory("Request headers")
		printKeybind("<Enter>", "Create new header")
		printKeybind("<Delete>", "Delete header in focus")
		printKeybind("<Tab>", "Focus in next input")
		printKeybind("<S-Tab>", "Focus in previous input")
		printKeybind("<S-ArrowRight>", "Go to next tab")
		printKeybind("<S-ArrowLeft>", "Go to previous tab")

		g.SetCurrentView("help_modal")
	}
	return nil
}
