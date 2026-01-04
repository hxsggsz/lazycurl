package collection

import (
	"fmt"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	"strings"

	"github.com/awesome-gocui/gocui"
)

type Header = map[string]string

var (
	activeHeaderIdx = 0
	totalHeaders    = 1
)

func Headers(g *gocui.Gui, maxX, maxY int) error {
	viewName := views.HEADERS
	x0, y0 := views.FULL, views.LAYOUT_INPUT_HEIGHT+views.LAYOUT_SECTION_Y_GAP
	x1, y1 := maxX/2, maxY-views.BOTTOM_MESSAGE-views.LOGS_BOTTOM

	if v, err := g.SetView(viewName, x0, y0, x1, y1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[3] Body *Headers"
	}

	for i := 0; i < totalHeaders; i++ {
		if err := renderHeaderPair(g, i, maxX); err != nil {
			return err
		}
	}

	return nil
}

func renderHeaderPair(g *gocui.Gui, index int, maxX int) error {
	rowY := views.LAYOUT_INPUT_HEIGHT + views.LAYOUT_SECTION_Y_GAP + 1 + (index * 3)

	keyName := fmt.Sprintf("header_key_%d", index)
	kx0, ky0, kx1, ky1 := 1, rowY, (maxX/4)-1, rowY+2

	if v, err := g.SetView(keyName, kx0, ky0, kx1, ky1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable, v.Title = true, "Key:"
		v.FrameColor = gocui.ColorYellow

		if err := setKeybindings(g, keyName); err != nil {
			return err
		}
	}

	valName := fmt.Sprintf("header_val_%d", index)
	vx0, vy0, vx1, vy1 := (maxX/4)+1, rowY, (maxX/2)-1, rowY+2

	if v, err := g.SetView(valName, vx0, vy0, vx1, vy1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable, v.Title = true, "Value:"

		if err := setKeybindings(g, valName); err != nil {
			return err
		}
	}

	return nil
}

func GetHeaders(g *gocui.Gui) Header {
	headers := make(Header, 0)

	for index := 0; index < totalHeaders; index++ {
		keyView := fmt.Sprintf("header_key_%d", index)
		valView := fmt.Sprintf("header_val_%d", index)

		if k, err := g.View(keyView); err == nil {
			headers[k.Buffer()] = ""
		}

		if v, err := g.View(valView); err == nil {
			headers[v.Buffer()] = v.Buffer()
		}
	}

	return headers
}

func createNewHeader(g *gocui.Gui, v *gocui.View) error {
	totalHeaders++
	activeHeaderIdx = totalHeaders - 1

	g.Update(func(g *gocui.Gui) error {
		newKeyView := fmt.Sprintf("header_key_%d", activeHeaderIdx)

		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		if nextV, err := g.SetCurrentView(newKeyView); err != nil {
			return err
		} else {
			nextV.FrameColor = gocui.ColorGreen
			nextV.TitleColor = gocui.ColorGreen
		}

		return nil
	})

	return nil
}

func focusNextInput(g *gocui.Gui, v *gocui.View) error {
	activeHeaderIdx = totalHeaders - 1

	g.Update(func(g *gocui.Gui) error {
		viewName := v.Name()
		v.FrameColor = gocui.ColorWhite
		v.TitleColor = gocui.ColorWhite

		if strings.HasPrefix(viewName, "header_key_") {
			v, _ := g.SetCurrentView(fmt.Sprintf("header_val_%d", activeHeaderIdx))
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen
		} else {
			v, _ := g.SetCurrentView(fmt.Sprintf("header_key_%d", activeHeaderIdx))
			v.FrameColor = gocui.ColorGreen
			v.TitleColor = gocui.ColorGreen
		}

		return nil
	})

	return nil
}

func deleteHeaderPair(g *gocui.Gui, v *gocui.View) error {
	g.Update(func(g *gocui.Gui) error {
		if activeHeaderIdx == 0 {
			return nil
		}

		totalHeaders--
		activeHeaderIdx = totalHeaders - 1

		oldKeyView := fmt.Sprintf("header_key_%d", activeHeaderIdx+1)
		oldValueView := fmt.Sprintf("header_val_%d", activeHeaderIdx+1)
		newKeyView := fmt.Sprintf("header_key_%d", activeHeaderIdx)

		g.DeleteView(oldKeyView)
		g.DeleteView(oldValueView)

		if nextV, err := g.SetCurrentView(newKeyView); err != nil {
			return err
		} else {
			nextV.FrameColor = gocui.ColorGreen
			nextV.TitleColor = gocui.ColorGreen
		}

		return nil
	})

	return nil
}

func setKeybindings(g *gocui.Gui, viewName string) error {

	headerKeyBindings := utils.KeybindsMaps{
		gocui.KeyArrowLeft:  {Modifier: gocui.ModShift, Handler: prevTab(BodyTabs)},
		gocui.KeyArrowRight: {Modifier: gocui.ModShift, Handler: nextTab(BodyTabs)},
		gocui.KeyEnter:      {Modifier: gocui.ModNone, Handler: createNewHeader},
		gocui.KeyTab:        {Modifier: gocui.ModNone, Handler: focusNextInput},
		gocui.KeyDelete:     {Modifier: gocui.ModNone, Handler: deleteHeaderPair},
	}

	views.HandleBlurInput(g, viewName)
	if err := utils.SetKeybind(g, headerKeyBindings, viewName); err != nil {
		return err
	}

	return nil
}
