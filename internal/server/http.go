package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/PeterBooker/dota2data/internal/files"
)

// startHTTP starts the HTTP server.
func (s *Server) startHTTP() {
	s.router = chi.NewRouter()

	// Middleware Stack
	//s.Router.Use(middleware.RequestID)
	//s.Router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.DefaultCompress)
	s.router.Use(middleware.RedirectSlashes)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	s.router.Use(middleware.Timeout(15 * time.Second))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	s.router.Use(cors.Handler)

	FileServer(s.router, "/assets")

	s.routes()

	s.serve()
}

func (s *Server) serve() {
	s.http = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      s.router,
		Addr:         ":" + s.Config.Port,
	}
	go func() { log.Fatal(s.http.ListenAndServe()) }()
}

func (s *Server) routes() {
	// Add Routes
	s.router.Get("/", s.index())
	s.router.Get("/docs", s.docs())
	s.router.Get("/privacy", s.privacy())

	// Need to disable RedirectSlashes middleware to enable this
	// redirects to /debug/prof/ which causes redirect loop
	//s.Router.Mount("/debug", middleware.Profiler())

	// Add API v1 routes
	s.router.Mount("/api/v1", s.apiRoutes())

	// Handle NotFound
	s.router.NotFound(s.notFound())
}

func (s *Server) apiRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/abilities", s.getAbilities())
	r.Get("/ability/{name}", s.getAbility())
	r.Get("/heroes", s.getHeroes())
	r.Get("/hero/{name}", s.getHero())
	r.Get("/items", s.getItems())
	r.Get("/item/{name}", s.getItem())
	r.Get("/units", s.getUnits())
	r.Get("/unit/{name}", s.getUnit())

	return r
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panicf("Failed to encode JSON: %v\n", err)
	}
}

func writeResp(w http.ResponseWriter, data interface{}) {
	writeJSON(w, data, http.StatusOK)
}

func writeError(w http.ResponseWriter, err error, status int) {
	writeJSON(w, map[string]string{
		"Error": err.Error(),
	}, status)
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.FileServer(files.Assets)

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=7776000")
		fs.ServeHTTP(w, r)
	}))
}