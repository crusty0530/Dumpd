package health

import (
	"fmt"
	"net/http"
)

// Status is
type Status struct {
	Message string
}

// ServeHTTP writes
func (h Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Message)
}
