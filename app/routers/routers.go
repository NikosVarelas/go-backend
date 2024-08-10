package routers

import (
	"go-backed/app/handlers"
	"net/http"
)

func InitRouters(r *http.ServeMux) {
	r.HandleFunc("/", handlers.HandleHome)
}

