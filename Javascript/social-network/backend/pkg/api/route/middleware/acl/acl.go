package acl

import (
	"context"
	"encoding/json"
	"net/http"

	"01.kood.tech/git/rols55/social-network/pkg/logger"
	"01.kood.tech/git/rols55/social-network/pkg/session"
)

type key string

const UserKey key = "user"

// DisallowAuth does not allow authenticated users to access the page
func DisallowAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := r.Context().Value(UserKey); user != nil {
			logger.Info("Unauthorized access for authenticated user")
			//http.NotFound(w, r)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "msg": "Unauthorized access or session has expired"}); err != nil {
				errStr := "Error encoding JSON"
				logger.Error(errStr)
				http.Error(w, errStr, http.StatusInternalServerError)
			}
			return
		}
		h.ServeHTTP(w, r)
	})
}

// DisallowAnon does not allow anonymous users to access the page
func DisallowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := r.Context().Value(UserKey); user == nil {
			logger.Info("Unauthorized access for anonymous user")
			//http.NotFound(w, r)
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]string{"status": "ERROR", "msg": "Unauthorized access or session has expired"}); err != nil {
				errStr := "Error encoding JSON"
				logger.Error(errStr)
				http.Error(w, errStr, http.StatusInternalServerError)
			}
			return
		}
		h.ServeHTTP(w, r)
	})
}

// Add user to the context if any
func AddUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sess, err := session.Get(r); err == nil {
			ctx := context.WithValue(r.Context(), UserKey, sess.UserId)
			h.ServeHTTP(w, r.WithContext(ctx))
			return
		} else if err != session.ErrExpired && err != session.ErrNotFound {
			logger.Error(err)
		}
		h.ServeHTTP(w, r)
	})
}
