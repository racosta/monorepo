// Package main contains the entry point for the HTTP server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	poker "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server"
)

// Repo-relative path to file
const dbFilename = "projects/go/learn_go_with_tests/http_server/cmd/webserver/game.db.json"

func mainNoExit() error {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFilename)
	defer closeFile()

	if err != nil {
		return fmt.Errorf("error creating player store: %v", err)
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
		return fmt.Errorf("could not listen on port 5000 %v", err)
	}

	return nil
}

func main() {
	if err := mainNoExit(); err != nil {
		log.Fatal(err)
	}
}
