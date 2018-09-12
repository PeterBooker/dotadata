package server

import (
	"context"
	"log"
	"net/http"

	"github.com/PeterBooker/dota2data/internal/config"
	"github.com/PeterBooker/dota2data/internal/data"
	"github.com/PeterBooker/dota2data/internal/jobs"
	"github.com/go-chi/chi"
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	Data   *data.Data
	router *chi.Mux
	http   *http.Server
}

// New returns a pointer to the main server struct
func New(log *log.Logger, config *config.Config) *Server {
	d := data.New()

	// Setup Server
	s := &Server{
		Config: config,
		Logger: log,
		Data:   d,
	}

	// Load Data
	s.Data.Load()
	//s.Data.Update()

	// Setup Cron
	jobs.Add("47 4 0 * * *", s.update)

	return s
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	s.startHTTP()
}

// Shutdown will release resources and stop the server.
func (s *Server) Shutdown(ctx context.Context) {
	s.http.Shutdown(ctx)
}

func (s *Server) update() {
	err := s.Data.Update()
	if err != nil {
		log.Printf("Update failed: %s\n", err)
	}
}