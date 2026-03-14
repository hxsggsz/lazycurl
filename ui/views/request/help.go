package request

import (
	"fmt"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"

	"github.com/awesome-gocui/gocui"
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

		if err := g.SetKeybinding("help_modal", gocui.KeyEsc, gocui.ModNone, helper.ToggleView("help_modal")); err != nil {
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
		v.Visible = false

		printCategory := func(name string) {

			title := views.CYAN + " -- " + name + " -- " + views.RESET
			title = utils.FormatLineFullWidth(v, fmt.Sprintf(title))
			fmt.Fprintf(v, "%s\n", title)
		}

		printKeybind := func(key, desc string) {
			keybind := views.GREEN + " " + key + "  " + views.RESET + desc
			keybind = utils.FormatLineFullWidth(v, fmt.Sprintf(keybind))
			fmt.Fprintf(v, "  %s\n", keybind)
		}

		printCategory("Global")
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

		printCategory("Request Files")
		printKeybind("<Tab>", "Focus the request files view")
		printKeybind("<C-/>", "Opens and close the request files view")
		printKeybind("<Space>", "Opens and close the request folder")
		printKeybind("<Enter>", "Opens and close the request folder")

		printCategory("Logs")
		printKeybind("<F10>", "Opens and close the logs view")

		g.SetCurrentView("help_modal")
	}
	return nil
}
