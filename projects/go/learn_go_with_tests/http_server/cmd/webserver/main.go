// Package main contains the entry point for the HTTP server.
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	poker "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server"
)

// Repo-relative path to file
const dbFilename = "projects/go/learn_go_with_tests/http_server/cmd/webserver/game.db.json"

func main() {
	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	playerServer := poker.NewPlayerServer(store)

	server := &http.Server{
		Addr:              ":5000",
		Handler:           playerServer,
		ReadHeaderTimeout: 5 * time.Second, // Mitigates Slowloris (G112)
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB limit to prevent OOM
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
