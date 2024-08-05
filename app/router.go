package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thanhhaudev/openapi-go/app/handler"
	"github.com/thanhhaudev/openapi-go/app/util"
)

func SetRoutes(r *mux.Router, h handler.AppHandler) {
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		util.Response(w, map[string]bool{"ok": true}, http.StatusOK)
	})

	a := s.PathPrefix("/auth").Subrouter()
	a.HandleFunc("/access", h.GetRefreshToken).Methods(http.MethodPost)
	a.HandleFunc("/exchange", h.GetAccessToken).Methods(http.MethodPost)
}
