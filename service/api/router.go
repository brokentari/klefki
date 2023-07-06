package api

import (
	"net/http"
	"time"

	"github.com/brokentari/klefki/service/v2/api/routes"
	"github.com/brokentari/klefki/service/v2/api/routes/auth"
	"github.com/brokentari/klefki/service/v2/api/routes/profile"
	"github.com/brokentari/klefki/service/v2/api/session"
	"github.com/brokentari/klefki/service/v2/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitRouter(root http.FileSystem, db *db.Database, sess *session.SessionManager) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept, Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false, 
		MaxAge: 300,
	}))
	

	authHandler := auth.AuthHandler{DB: db, Session: sess}
	profileHandler := profile.ProfileHandler{DB: db, Session: sess}

	r.Group(func(r chi.Router) {
		r.Get("/health_check", routes.HealthCheckHandler)
		r.Post("/register", authHandler.RegisterHandler)
		r.Post("/login", authHandler.LoginHandler)
		r.Get("/profile", profileHandler.HandleGetProfile)

	})

	r.Group(func(r chi.Router) {
		r.Use(KlefkiAuth(db))
	})

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	// r.Get("/", fileServer(root))

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

// func fileServer(root http.FileSystem) http.HandlerFunc   {
// 	// if strings.Contains(path, "{}*") {
// 	// 	panic("fileserver does not permit any url parameters")
// 	// }

// 	// if path != "/" && path[len(path)-1] != '/' {
// 	// 	r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
// 	// 	path += "/"
// 	// }
// 	// path += "*"

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		rctx := chi.RouteContext(r.Context())
// 		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
// 		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
// 		fs.ServeHTTP(w, r)
// 	}
// }