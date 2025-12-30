package ui

import (
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	typedValue, err := views.Input(g, maxX)
	if err != nil {
		return err
	}

	if typedValue != "" {
		// essa funćão é chamada toda hora então usar os logs apenas em ações do usuário para o notifica-lo uma vez
		log.Println("digitando ->", typedValue)
	}

	if err := views.Output(g, maxX, maxY); err != nil {
		return err
	}

	if err := views.Body(g, maxX, maxY); err != nil {
		return err
	}

	RegisterGlobalNumericNavigation(g)
	return nil
}
