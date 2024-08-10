package handlers

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "app/views/index.html")
}
