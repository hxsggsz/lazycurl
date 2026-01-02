package ui

import (
	"lazycurl/ui/keyboard"
	"lazycurl/ui/views"
	"lazycurl/ui/views/collection"
	"log"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := collection.Method(g, maxX, maxY); err != nil {
		return err
	}

	typedValue, err := collection.Input(g, maxX)
	if err != nil {
		return err
	}

	if typedValue != "" {
		// essa funćão é chamada toda hora então usar os logs apenas em ações do usuário para o notifica-lo uma vez
		log.Println("digitando ->", typedValue)
	}

	// if err := collection.Tabs(g, maxX, maxY); err != nil {
	// 	return err
	// }

	if err := collection.Headers(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Body(g, maxX, maxY); err != nil {
		return err
	}

	if err := collection.Response(g, maxX, maxY); err != nil {
		return err
	}

	if err := views.Output(g, maxX, maxY); err != nil {
		return err
	}

	keyboard.RegisterGlobalNumericNavigation(g)
	keyboard.RegisterGlobalSubmit(g)
	keyboard.RegisterGlobalViewNavigation(g)

	return nil
}
