package keyboard

import (
	"fmt"
	"lazycurl/pkg/highlight"
	"lazycurl/pkg/request"
	"lazycurl/ui/utils"
	"lazycurl/ui/views"
	requestView "lazycurl/ui/views/request"
	"log"
	"regexp"
	"strings"

	"github.com/awesome-gocui/gocui"
)

var (
	responseChan = make(chan string, 1)
	headerChan   = make(chan map[string]string, 1)
)

func RegisterGlobalSubmit(g *gocui.Gui) error {
	viewsToSubmitRequest := []string{views.URL, views.RESPONSE, views.RESPONSE_HEADERS, views.METHOD}

	for _, name := range viewsToSubmitRequest {
		if err := g.SetKeybinding(name, gocui.KeyEnter, gocui.ModNone, submitHandler); err != nil {
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

	go func() {
		for headers := range headerChan {
			var allHeaders strings.Builder
			for key, value := range headers {
				fmt.Fprintf(&allHeaders, " %s: %s\n", key, value)
			}

			finalContent := allHeaders.String()
			totalHeaders := len(headers)

			g.Update(func(g *gocui.Gui) error {
				return UpdateHeadersView(g, finalContent, totalHeaders)
			})
		}
	}()

	return nil
}

func submitHandler(g *gocui.Gui, v *gocui.View) error {
	responseChan <- "loading..."

	go func() {
		log.Println("submitting request...")

		res := request.RequestBuilder(
			requestView.GetCurrentMethod(g),
			requestView.GetInputValue(g),
			requestView.GetBodyValue(g),
			requestView.GetHeaders(g),
		).Send()

		statusMsg := coloredStatus(res.StatusCode)
		responseChan <- fmt.Sprintf("status: %s \n\n\n %s", statusMsg, res.Body)
		headerChan <- res.Headers
	}()

	return nil
}

func UpdateHeadersView(g *gocui.Gui, content string, totalHeaders int) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(views.RESPONSE_HEADERS)
		if err != nil {
			return err
		}

		v.Clear()
		v.Title = fmt.Sprintf("[4] Response *Headers (%d)", totalHeaders)

		log.Println("content", content)
		fmt.Fprint(v, utils.FormatLineFullWidth(v, content))

		resView, err := g.View(views.RESPONSE)
		if err != nil {
			return err
		}
		resView.Title = fmt.Sprintf("[4] *Response Headers (%d)", totalHeaders)

		return nil
	})
	return nil
}
func UpdateResponseView(g *gocui.Gui, content string) error {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View(views.RESPONSE)
		if err != nil {
			return err
		}

		v.Clear()

		parts := strings.SplitN(content, "\n", 2)
		log.Println("parts", parts)

		if len(parts) >= 2 {
			highlighetdContent := highlight.Json(parts[1])
			fmt.Fprint(v, utils.FormatLineFullWidth(v, parts[0]))
			fmt.Fprint(v, utils.FormatLineFullWidth(v, highlighetdContent))

			return nil
		}

		highlighetdContent := highlight.Json(content)

		fmt.Fprint(v, utils.FormatLineFullWidth(v, highlighetdContent))

		return nil
	})
	return nil
}

func coloredStatus(statusCode int) string {
	statusStr := fmt.Sprintf("status: %d", statusCode)
	numStr := fmt.Sprintf("%d", statusCode)

	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	cleanStatus := re.ReplaceAllString(statusStr, "")
	cleanNum := re.ReplaceAllString(numStr, "")

	firstDigit := statusCode / 100

	var color string
	switch firstDigit {
	case 2:
		color = views.GREEN
	case 4:
		color = views.YELLOW
	case 5:
		color = views.RED
	default:
		color = ""
	}

	idx := strings.Index(statusStr, numStr)
	return color + cleanNum + views.RESET + cleanStatus[idx+len(cleanNum):]
}
