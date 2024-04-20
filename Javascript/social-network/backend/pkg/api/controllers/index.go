package controllers

import (
	"fmt"
	"net/http"
)

// Home/Index page handler/controller
func (h *BaseController) Index(w http.ResponseWriter, r *http.Request) {
	resp := []byte(`{"status": "ok"}`)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp)))
	w.Write(resp)
}
