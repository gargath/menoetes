package modules_api

import (
	m "github.com/gargath/menoetes/pkg/server/middleware"
	"github.com/gargath/menoetes/pkg/store"
	"github.com/gorilla/mux"
)

func RegisterModulesAPI(r *mux.Router) {
	api := &modulesAPI{store: store.NewMockStore()}
	s := r.PathPrefix("/v1/modules").Subrouter()
	s.HandleFunc("", m.Use(api.listHandler, m.TokenAuth)).Methods("GET")
	s.HandleFunc("/", m.Use(api.listHandler, m.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}", m.Use(api.listHandler, m.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}", m.Use(api.listHandler, m.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/versions", m.Use(api.versionsHandler, m.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}/{provider}/{version}", m.Use(api.versionDetailsHandler, m.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}/{provider}", m.Use(api.versionDetailsHandler, m.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/download", m.Use(api.latestDownloadHandler, m.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/{version}/download", m.Use(api.downloadHandler, m.TokenAuth)).Methods("GET")
}
