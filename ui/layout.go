package ui

import (
	"lazycurl/ui/components"
	"log"

	"github.com/awesome-gocui/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	value, err := components.Input(g, maxX)
	if err != nil {
		return err
	}

	// essa funćão é chamada toda hora então usar os logs apenas em aćoes do usuário para o notifica-lo uma vez
	log.Println("digitado ->", value)

	if err := components.Output(g, maxX, maxY); err != nil {
		return err
	}

	return nil
}
