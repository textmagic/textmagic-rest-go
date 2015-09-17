package textmagic

import (
	"errors"
	"fmt"
)

// ErrPing represents a ping error.
var ErrPing = errors.New("unable to ping API")

// Error represents a TextMagic API error.
type Error struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}

// NewError returns a new Error with the
// given code number and message string.
func NewError(c int, m string) *Error {
	return &Error{Code: c, Message: m}
}

// Error implements the error interface for
// the Error struct.
func (e *Error) Error() string {
	var message = "TextMagic Rest API Error:"

	if e.Code != 0 {
		message += fmt.Sprintf(" Code: %d", e.Code)
	} else if e.Message != "" {
		message += fmt.Sprintf(", Message: %s", e.Message)
	} else if len(e.Errors) > 0 {
		message += fmt.Sprintf(", Errors: %s", e.Errors)
	}

	return message
}
