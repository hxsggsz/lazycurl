package collection

import (
	"fmt"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"

	"github.com/awesome-gocui/gocui"
)

func inputViewConfig(g *gocui.Gui, v *gocui.View) {
	v.Title = fmt.Sprintf("[%d] %s", 2, utils.Capitalize(views.URL))
	v.Editable = true
	g.Cursor = true

}

func Input(g *gocui.Gui, maxX int) error {
	viewName := views.URL
	x0 := 10
	y0 := views.FULL

	x1 := maxX - views.LAYOUT_SECTION_Y_GAP
	y1 := views.LAYOUT_INPUT_HEIGHT

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// default focus when iniciate the view
		g.SetCurrentView(viewName)
		inputViewConfig(g, v)
		views.HandleBlurInput(g, viewName)

		if hasFocus := utils.ViewHasFocus(g, viewName); hasFocus {
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen
		}
	}

	return nil
}

func GetInputValue(g *gocui.Gui) string {
	v, err := g.View(views.URL)
	if err != nil {
		return ""
	}
	return v.Buffer()
}
