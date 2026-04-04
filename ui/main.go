package ui

import (
	"lazycurl/output"
	"lazycurl/pkg/collection"
	"lazycurl/ui/options"
	"log"

	"github.com/awesome-gocui/gocui"
)

type AddFoldersFunc func(foldersPath string) error

func InitLayout(collectionPath string) {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Mouse = true

	log.SetOutput(&output.LogViewWriter{Gui: g})
	clt := collection.NewCollection(collectionPath)
	clt.LoadCollectionFiles()

	if len(clt.Files) > 0 {
		g.SetManagerFunc(layout(clt))
	}

	options.QuitKeyByind(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
