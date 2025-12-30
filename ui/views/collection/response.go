package collection

import (
	"fmt"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

const (
	LayoutInputHeight = 1
	LayoutSectionGap  = 2
	LogsHeight        = 3
)

const (
	InputTop    = 0
	InputHeight = LayoutInputHeight
	BodyTop     = LayoutInputHeight + LayoutSectionGap
)

func Response(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.RESPONSE
	height := maxY - LogsHeight // reservar espaço para logs

	// Largura: 50% da tela, posicionada à DIREITA
	width := maxX / 2
	x0 := width + LayoutSectionGap
	x1 := maxX - 1

	y0 := BodyTop // Sem magic number!
	y1 := height - 1

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf("[%d] %s", 3, utils.Capitalize(viewName))
		v.Autoscroll = true
		v.Wrap = true
		g.Cursor = true
		g.SetViewOnBottom(viewName)
	}
	return nil
}
