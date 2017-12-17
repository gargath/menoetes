package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type registryServer struct{
	tlsCertFilePath string
	tlsKeyFilePath string
}

type helloHandler struct{}

func New(cert string, key string) *registryServer {
	return &registryServer {
		tlsCertFilePath: cert,
		tlsKeyFilePath: key,
	}
}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request Headers:\n")
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			log.Printf("%v: %v", name, h)
   		}
 	}
	if (r.Header.Get("Authorization") != "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders") {
		log.Println("Unauthorized Access")
		http.Error(w, "access denied", http.StatusForbidden)
	}
        if (r.URL.Path == "/.well-known/terraform.json") {
		log.Println("Service Discovery Request")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\n\"modules.v1\": \"https://modules.howthe.click/v1/\"\n}")
	} else {
		fmt.Fprintf(w, "hello, you've hit %s\n", r.URL.Path)
        	log.Printf("Request for %s\n", r.URL.Path)
	}
}


func (s *registryServer) Run() {
	err := http.ListenAndServeTLS(":443", s.tlsCertFilePath, s.tlsKeyFilePath, helloHandler{})
	log.Fatal(err)
}
