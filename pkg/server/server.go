package server

import (
	"net/http"

	modules_api "github.com/gargath/menoetes/pkg/modules_api"
	m "github.com/gargath/menoetes/pkg/server/middleware"
	s "github.com/gargath/menoetes/pkg/store"

	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type registryServer struct {
	debug           bool
	tlsCertFilePath string
	tlsKeyFilePath  string
	store           s.Store
}

func New(cert string, key string, d bool, st s.Store) *registryServer {
	return &registryServer{
		debug:           d,
		tlsCertFilePath: cert,
		tlsKeyFilePath:  key,
		store:           st,
	}
}

func (s *registryServer) Run() {
	r := mux.NewRouter()
	if s.debug {
		r.HandleFunc("/.well-known/{disco_path}", m.Use(discoHandler, m.TokenAuth, m.DumpHeaders))
	} else {
		r.HandleFunc("/.well-known/{disco_path}", m.Use(discoHandler, m.TokenAuth))
	}
	modules_api.RegisterModulesAPI(r)
	log.Info("Listening...")
	err := http.ListenAndServeTLS(":443", s.tlsCertFilePath, s.tlsKeyFilePath, r)
	log.Fatal(err)
}
