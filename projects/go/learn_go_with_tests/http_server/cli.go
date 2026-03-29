package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI is a command-line interface for the poker game that interacts with the PlayerStore to record wins.
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI creates a new CLI instance with the given PlayerStore and input reader.
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in)}
}

// PlayPoker simulates playing a game of poker.
func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
