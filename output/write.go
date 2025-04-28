package output

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

type LogViewWriter struct {
	Gui *gocui.Gui
}

func (w *LogViewWriter) Write(p []byte) (int, error) {
	w.Gui.Update(func(g *gocui.Gui) error {
		v, err := g.View("logs")
		if err != nil {
			return err
		}

		fmt.Fprint(v, string(p))
		return nil
	})

	return len(p), nil
}
