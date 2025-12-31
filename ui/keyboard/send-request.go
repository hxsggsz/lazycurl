package keyboard

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

func RegisterGlobalSubmit(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, submitHandler()); err != nil {
		return err
	}

	return nil
}

func submitHandler() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		log.Printf("submitting request...")

		res := request.RequestBuilder().
			SetMethod(request.GET).
			SetURL("https://jsonplaceholder.typicode.com/posts").
			Build()

		log.Printf("successfully submitted request. \n")
		UpdateResponseView(g, res.Body)
		return nil
	}
}

func UpdateResponseView(g *gocui.Gui, content string) error {
	v, err := g.View(views.RESPONSE)
	if err != nil {
		return err
	}

	// Limpa o conteúdo anterior
	v.Clear()

	// Escreve o novo conteúdo
	fmt.Fprint(v, content)

	return nil
}
