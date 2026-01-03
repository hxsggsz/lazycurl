package keyboard

import (
	"fmt"
	"lazycurl/pkg/request"
	"lazycurl/ui/views"
	"log"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m" // 5xx
	Yellow = "\033[33m" // 4xx
	Green  = "\033[32m" // 2xx

	responseChan = make(chan string, 1)
)

func RegisterGlobalSubmit(g *gocui.Gui) error {
	viewsToSubmitRequest := []string{views.URL, views.RESPONSE, views.METHOD}

	for _, name := range viewsToSubmitRequest {
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
				SetURL("https://jsonplaceholder.typicode.com/post").
				Build()

			statusMsg := coloredStatus(res.StatusCode)
			responseChan <- fmt.Sprintf("status: %s \n \n %s", statusMsg, res.Body)
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

func coloredStatus(statusCode int) string {
	statusStr := fmt.Sprintf("status: %d", statusCode)
	numStr := fmt.Sprintf("%d", statusCode)

	firstDigit := statusCode / 100 // 2, 4 ou 5

	var color string
	switch {
	case firstDigit == 2:
		color = Green
	case firstDigit == 4:
		color = Yellow
	case firstDigit == 5:
		color = Red
	default:
		color = "" // Sem cor para outros (1xx, 3xx)
	}

	// Substitui o número pela versão colorida
	idx := strings.Index(statusStr, numStr)
	return color + numStr + Reset + statusStr[idx+len(numStr):]
}
