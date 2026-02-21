package ui

import (
	"lazycurl/output"
	"lazycurl/ui/options"
	"log"

	"github.com/awesome-gocui/gocui"
)

func InitLayout(collectionPath string) {
	log.Println("Initializing UI with collection:", collectionPath)
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	log.SetOutput(&output.LogViewWriter{Gui: g})

	options.QuitKeyByind(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
