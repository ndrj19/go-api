package handlers

import (
	"net/http"
)

func Landing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API landing page!"))
}
