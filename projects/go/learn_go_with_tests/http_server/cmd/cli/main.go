// Package main is the entry point for our poker CLI application.
package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server"
)

// Repo-relative path to file
const dbFilename = "projects/go/learn_go_with_tests/http_server/cmd/cli/game.db.json"

func mainNoExit() error {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFilename)
	defer closeFile()

	if err != nil {
		return fmt.Errorf("error creating player store: %v", err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()

	return nil
}

func main() {
	if err := mainNoExit(); err != nil {
		log.Fatal(err)
	}
}
