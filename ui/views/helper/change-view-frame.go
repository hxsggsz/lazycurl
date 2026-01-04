package helper

import "github.com/awesome-gocui/gocui"

func ChangeViewFrame(g *gocui.Gui) {
	views := g.Views()

	for _, view := range views {
		view.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
	}
}
