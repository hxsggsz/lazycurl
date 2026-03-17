package filetree

import (
	"lazycurl/ui/views"
	"lazycurl/ui/views/helper"
	"log"
	"time"

	"github.com/awesome-gocui/gocui"
)

func addFolder(g *gocui.Gui, maxX, maxY int, addFolderFunc func(foldersPath string) error) error {
	x0 := maxX / 4
	x1 := (maxX / 4) * 3

	y0 := maxY / 5
	y1 := y0 + 2

	v, err := g.SetView(views.ADD_FOLDER, x0, y0, x1, y1, 0)
	if err != nil {
		if err != gocui.ErrUnknownView {
			log.Println("Erro ao criar modal:", err)
			return err
		}

		// Configurações visuais do Modal
		v.Title = " Folder(s) Name(s) "
		v.Subtitle = " Separate multiple folders with slashes (e.g., folder1/folder2) "
		v.Frame = true
		v.Wrap = true
		v.Editable = true
		v.Visible = true

		v.FrameColor = gocui.ColorGreen
		v.TitleColor = gocui.ColorGreen

	}

	g.SetKeybinding(views.ADD_FOLDER, gocui.KeyEnter, gocui.ModNone, createFolders(addFolderFunc))
	return nil
}

func getFoldersInput(g *gocui.Gui) string {
	v, err := g.View(views.ADD_FOLDER)
	if err != nil {
		log.Println("Erro ao obter view do modal:", err)
		return ""
	}
	return v.Buffer()
}

func createFolders(addFolderFunc func(foldersPath string) error) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		foldersPath := getFoldersInput(g)
		if foldersPath == "" {
			return nil // Ou exibir uma mensagem de erro
		}

		if err := addFolderFunc(foldersPath); err != nil {
			log.Println("Erro ao criar pasta:", err)
			return err
		}

		log.Printf("Pasta(s) '%s' criada(s) com sucesso!", foldersPath)
		helper.CloseView(views.ADD_FOLDER)(g, v)
		views.ShowToast(g, "folder created successfully", "success", 2*time.Second)
		return nil
	}
}

// func setupModalKeys(g *gocui.Gui) error {
// 	g.SetKeybinding(views.ADD_FOLDER, gocui.KeyEnter, gocui.ModNone, createFolders(addFolderFunc))

// 	return nil
// }
