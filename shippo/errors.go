package shippo

import (
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response
}

func (e *ErrorResponse) Error() string {
	b, _ := io.ReadAll(e.Response.Body)
	return fmt.Sprintf("%v %v: %d %+v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, string(b))
}
