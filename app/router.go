package app

import (
	"github.com/gorilla/mux"
	"github.com/thanhhaudev/openapi-go/app/handler"
	"net/http"
)

func SetRoutes(r *mux.Router, h handler.AppHandler) {
	r.HandleFunc("/api/health", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Gorilla!\n"))
	})
}
