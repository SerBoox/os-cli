package exceptions

import (
	"fmt"
)

// JSONError declare exception structure
type JSONError struct {
	Method  string
	Message string
	JSON    string
}

func (err *JSONError) Error() string {
	return fmt.Sprintf(
		"JSONError exception: %s: %s\n %s",
		err.Method, err.Message, err.JSON,
	)
}
