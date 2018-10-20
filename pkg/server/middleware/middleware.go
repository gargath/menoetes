package middleware

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func TokenAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer reallylongstringthatstotallygoingtostandoutinthelistofheaders" {
			log.Printf("Unauthorized Access: %+v", *r)
			http.Error(w, "access denied", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	}
}

func DumpHeaders(h http.HandlerFunc) http.HandlerFunc {
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
