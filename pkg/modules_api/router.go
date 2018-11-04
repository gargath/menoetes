package modules_api

import (
	m "github.com/gargath/menoetes/pkg/server/middleware"
	store "github.com/gargath/menoetes/pkg/store"
	"github.com/gorilla/mux"
)

func RegisterModulesAPI(r *mux.Router, tm *m.TokenManager, store store.Store) {
	api := &modulesAPI{store: store}
	s := r.PathPrefix("/v1/modules").Subrouter()
	s.HandleFunc("", m.Use(api.listHandler, tm.TokenAuth)).Methods("GET")
	s.HandleFunc("/", m.Use(api.listHandler, tm.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}", m.Use(api.listHandler, tm.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}", m.Use(api.listHandler, tm.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/versions", m.Use(api.versionsHandler, tm.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}/{provider}/{version}", m.Use(api.versionDetailsHandler, tm.TokenAuth)).Methods("GET")
	s.HandleFunc("/{namespace}/{name}/{provider}", m.Use(api.versionDetailsHandler, tm.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/download", m.Use(api.latestDownloadHandler, tm.TokenAuth)).Methods("GET")

	s.HandleFunc("/{namespace}/{name}/{provider}/{version}/download", m.Use(api.downloadHandler, tm.TokenAuth)).Methods("GET")
}
