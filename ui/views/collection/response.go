package collection

import (
	"lazycurl/ui/utils"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func Response(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.RESPONSE
	height := maxY - views.BOTTOM_MESSAGE

	width := maxX / 2
	x0 := width + views.LAYOUT_SECTION_X_GAP
	x1 := maxX - views.RIGHT_BORDER

	y0 := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
	y1 := height - views.LOGS_BOTTOM

	responseHeaders(g, maxX, maxY)

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "[4] *Response Headers"
		v.Autoscroll = false
		v.Wrap = true
		v.HasLoader = true
		g.Cursor = true
	}

	headerKeyBindings := utils.KeybindsMaps{
		gocui.KeyArrowLeft:  {Modifier: gocui.ModShift, Handler: prevTab(ResponseTabs)},
		gocui.KeyArrowRight: {Modifier: gocui.ModShift, Handler: nextTab(ResponseTabs)},
	}

	if err := utils.SetKeybind(g, headerKeyBindings, viewName); err != nil {
		return err
	}

	return nil
}

func responseHeaders(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.RESPONSE_HEADERS
	height := maxY - views.BOTTOM_MESSAGE

	width := maxX / 2
	x0 := width + views.LAYOUT_SECTION_X_GAP
	x1 := maxX - views.RIGHT_BORDER

	y0 := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP
	y1 := height - views.LOGS_BOTTOM

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[4] Response *Headers"
		v.Wrap = true
	}

	headerKeyBindings := utils.KeybindsMaps{
		gocui.KeyArrowLeft:  {Modifier: gocui.ModShift, Handler: prevTab(ResponseTabs)},
		gocui.KeyArrowRight: {Modifier: gocui.ModShift, Handler: nextTab(ResponseTabs)},
	}

	if err := utils.SetKeybind(g, headerKeyBindings, viewName); err != nil {
		return err
	}

	return nil
}
