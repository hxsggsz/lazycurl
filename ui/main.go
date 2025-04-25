package ui

import (
	"lazycurl/ui/options"
	"log"

	"github.com/awesome-gocui/gocui"
)

func InitLayout() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	options.QuitKeyByind(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
