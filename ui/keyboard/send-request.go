package keyboard

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/views"
	"log"

	"github.com/awesome-gocui/gocui"
)

var responseChan = make(chan string, 1)

func RegisterGlobalSubmit(g *gocui.Gui) error {
	viewsWithEnter := []string{views.URL, views.RESPONSE, views.METHOD}

	for _, name := range viewsWithEnter {
		if err := g.SetKeybinding(name, gocui.KeyEnter, gocui.ModNone, submitHandler()); err != nil {
			return err
		}
	}

	go func() {
		for content := range responseChan {
			g.Update(func(g *gocui.Gui) error {
				return UpdateResponseView(g, content)
			})
		}
	}()

	return nil
}

func submitHandler() func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		responseChan <- "loading..."

		go func() {
			log.Println("submitting request...")

			res := request.RequestBuilder().
				SetMethod(request.GET).
				SetURL("https://jsonplaceholder.typicode.com/posts").
				Build()

			log.Println("success!")
			responseChan <- fmt.Sprintf("status: %d\n \n %s", res.StatusCode, res.Body)
		}()

		return nil
	}
}

func UpdateResponseView(g *gocui.Gui, content string) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(views.RESPONSE)
		if err != nil {
			return err
		}

		v.Clear()

		fmt.Fprint(v, content)

		return nil
	})
	return nil
}
