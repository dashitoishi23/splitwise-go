package util

import (
	"net/http"

	commonmodels "github.com/dashitoishi23/splitwise-go/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RootHttpHandler(servers []commonmodels.HttpServerConfig) http.Handler {
	r := mux.NewRouter()

	for _, server := range servers {
		r.Handle(server.Route, server.Server).Methods(server.Methods...)
	}

	var handler http.Handler

	handler = r

	handler = cors.AllowAll().Handler(r)

	return handler
}
