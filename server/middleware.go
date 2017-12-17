package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
  "strings"
	)

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func tokenAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Header.Get("Authorization") != "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders") {
			log.Println("Unauthorized Access")
			http.Error(w, "access denied", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func dumpHeaders(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				log.WithFields(log.Fields{name: h}).Debug("Request Headers:")
			}
		}
		h.ServeHTTP(w, r)
	}
}
