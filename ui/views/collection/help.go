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

		printCategory := func(name string) {
			fmt.Fprintf(v, "%s  -- %s --%s\n", ColorBlue, name, ColorReset)
		}

		printKeybind := func(key, desc string) {
			fmt.Fprintf(v, "  %s%-10s%s %s\n", ColorGreen, key, ColorReset, desc)
		}

		// --- Conteúdo Baseado no Vídeo ---
		printCategory("Global")
		printKeybind("<enter>", "Submit Request")
		printKeybind("<tab>", "Next Field")
		printKeybind("<f10>", "Toggle Logs")
		printKeybind("q", "Quit Help")

		printCategory("Navigation")
		printKeybind("<up>", "Scroll up")
		printKeybind("<down>", "Scroll down")
		g.SetCurrentView("help_modal")
	}
	return nil
}
