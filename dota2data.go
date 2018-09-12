package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/PeterBooker/dota2data/internal/config"
	"github.com/PeterBooker/dota2data/internal/db"
	"github.com/PeterBooker/dota2data/internal/log"
	"github.com/PeterBooker/dota2data/internal/server"
	"github.com/PeterBooker/dota2data/internal/jobs"
)

var (
	version string
	commit  string
	date    string
)

//go:generate go run -tags=dev embed_files.go

func main() {
	// Set and Parse flags
	port := flag.String("port", "9001", "Port to serve HTTP requests on.")
	flag.Parse()

	// Create Logger
	l := log.New()

	l.Println("Starting Dota2Data...")

	// Create Config
	c := config.Setup(version, commit, date, *port)

	// Setup BoltDB
	db.Setup(c.WD)
	defer db.Close()

	// Setup server struct to hold all App data
	s := server.New(l, c)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Setup HTTP server.
	s.Setup()

	// Start Job Runner
	jobs.Start()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Shutdown
	s.Shutdown(ctx)
	jobs.Stop()

}