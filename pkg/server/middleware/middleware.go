package middleware

import (
	ctx "context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gargath/menoetes/pkg/store"
	log "github.com/sirupsen/logrus"
)

type TokenManager struct {
	Store store.Store
}

func Use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}

func (tm *TokenManager) TokenAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "anonymous access not allowed", http.StatusUnauthorized)
			return
		}
		tns := strings.Split(auth, " ")
		fmt.Printf(">>%s<<\n", tns[0])
		if strings.TrimSpace(tns[0]) != "Bearer" {
			fmt.Printf("Bearer isn't")
			http.Error(w, "malformed Authorization header", http.StatusBadRequest)
			return
		}
		user, err := tm.Store.ValidateAccessToken(strings.TrimSpace(tns[1]))
		if err != nil {
			log.Printf("Unauthorized Access: %+v", *r)
			http.Error(w, "access denied", http.StatusForbidden)
			return
		} else {
			ctx := ctx.WithValue(r.Context(), "username", user)
			r2 := r.WithContext(ctx)
			h.ServeHTTP(w, r2)
		}
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
