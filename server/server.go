package server

import (
	log "github.com/sirupsen/logrus"
	mux "github.com/gorilla/mux"
	"net/http"
)

type registryServer struct{
	debug bool
	tlsCertFilePath string
	tlsKeyFilePath string
}

func New(cert string, key string, d bool) *registryServer {
	return &registryServer {
		debug: d,
		tlsCertFilePath: cert,
		tlsKeyFilePath: key,
	}
}



func (s *registryServer) Run() {
	r := mux.NewRouter()
	if (s.debug) {
		r.HandleFunc("/.well-known/{disco_path}", use(discoHandler, tokenAuth, dumpHeaders))
	} else {
		r.HandleFunc("/.well-known/", use(discoHandler, tokenAuth))
	}
	registerModulesAPI(r)
	log.Info("Listening...")
	err := http.ListenAndServeTLS(":443", s.tlsCertFilePath, s.tlsKeyFilePath, r)
	log.Fatal(err)
}
