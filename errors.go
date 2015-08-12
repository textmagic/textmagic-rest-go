package textmagic

import (
	"fmt"
)

type TextmagicError struct {
	Code    uint16                 `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

func (e TextmagicError) Error() string {
	var message = "TextMagic Rest API Error:"

	if e.Code != 0 {
		message = fmt.Sprintf("%s Code: %d", message, e.Code)
	}
	if e.Message != "" {
		message = fmt.Sprintf("%s, Message: %s", message, e.Message)
	}
	if e.Errors != nil {
		message = fmt.Sprintf("%s, Errors: %s", message, e.Errors)
	}

	return message
}
