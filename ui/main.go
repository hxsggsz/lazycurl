package ui

import (
	"lazycurl/output"
	"lazycurl/pkg/collection"
	"lazycurl/ui/options"
	"log"

	"github.com/awesome-gocui/gocui"
)

func InitLayout(collectionPath string) {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	clt := collection.NewCollection(collectionPath)
	clt.LoadCollectionFiles()

	g.SetManagerFunc(layout(clt.Files))

	log.SetOutput(&output.LogViewWriter{Gui: g})

	options.QuitKeyByind(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
