package views

const (
	METHOD           = "method"
	URL              = "url"
	TABS             = "tabs"
	BODY             = "body"
	HEADERS          = "headers"
	RESPONSE         = "response"
	RESPONSE_HEADERS = "response_headers"
	LOGS             = "logs"
)

const (
	FULL                 = 0
	LOGS_BOTTOM          = 1
	BOTTOM_MESSAGE       = 1
	LAYOUT_INPUT_HEIGHT  = 2
	LAYOUT_SECTION_X_GAP = 2
	LAYOUT_SECTION_Y_GAP = 1
	RIGHT_BORDER         = 1 // only right views need this border margin
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	YELLOW = "\033[33m"
	GREEN  = "\033[32m"
)
