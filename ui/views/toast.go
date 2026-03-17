package views

import (
	"fmt"
	"time"

	"github.com/awesome-gocui/gocui"
)

func ShowToast(g *gocui.Gui, message string, variant string, duration time.Duration) {
	viewName := "toast"
	maxX, maxY := g.Size()

	width := len(message) + 4
	x0 := (maxX / 2) - (width / 2)
	x1 := x0 + width
	y0 := maxY - 5
	y1 := y0 + 2

	v, err := g.SetView(viewName, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
		v.Frame = true
		v.FrameColor = setVariant(variant)
		v.BgColor = gocui.ColorBlack
		v.FgColor = gocui.ColorWhite | gocui.AttrBold
		fmt.Fprintf(v, " %s ", message)
	}

	go func() {
		time.Sleep(duration)

		g.Update(func(g *gocui.Gui) error {
			return g.DeleteView(viewName)
		})
	}()
}

func setVariant(variant string) gocui.Attribute {
	switch variant {
	case "success":
		return gocui.ColorGreen
	case "error":
		return gocui.ColorRed
	case "info":
		return gocui.ColorBlue
	default:
		return gocui.ColorCyan
	}
}
