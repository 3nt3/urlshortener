package handlers

import (
	"fmt"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "irgendwas 10.0\n")
}
