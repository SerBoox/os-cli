package exceptions

import (
	"fmt"
)

// HTTPError declare exception structure
type HTTPError struct {
	Method  string
	Message string
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf(
		"HTTPError exception: %s: %s",
		err.Method, err.Message,
	)
}
