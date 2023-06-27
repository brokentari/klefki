package api

import (
	"net/http"
	"time"

	"github.com/brokentari/klefki/service/v2/api/routes/auth"
	"github.com/brokentari/klefki/service/v2/api/routes/profile"
	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(db *db.Database, sess *session.SessionManager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	

	authHandler := auth.AuthHandler{DB: db, Session: sess}
	profileHandler := profile.ProfileHandler{DB: db, Session: sess}

	r.Group(func(r chi.Router) {
		r.Post("/register", authHandler.RegisterHandler)
		r.Get("/login", authHandler.LoginHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(KlefkiAuth(db))
		r.Get("/profile", profileHandler.HandleGetProfile)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}

func KlefkiAuth(db *db.Database) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")
			if authToken != "" {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		})
	}
}
