package filetree

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

func AddFile(g *gocui.Gui, maxX, maxY int) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	// --- CÁLCULO VERTICAL (Ao topo) ---
	// y0: 20% de distância do topo da tela
	// y1: y0 + 2 (tamanho suficiente para uma linha de input + bordas)
	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView("add-file-modal", x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		// Configurações visuais do Modal
		v.Title = " Input "
		v.Frame = true
		v.Wrap = true
		v.Editable = true // Se for para entrada de texto
		v.Visible = true

		// Cores para destacar que é um modal
		v.FrameColor = gocui.ColorCyan
		v.TitleColor = gocui.ColorCyan

	}
	g.SetCurrentView("add-file-modal")
	g.SetViewOnTop("add-file-modal")
	return nil
}
