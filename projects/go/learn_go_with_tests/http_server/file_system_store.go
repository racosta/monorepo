package poker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

// FileSystemPlayerStore implements PlayerStore and persists data in the filesystem.
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   leagueLib.League
}

// NewFileSystemPlayerStore initializes a FileSystemPlayerStore with the given file path.
// If the file does not exist, it will be created. If the file exists but is empty, it will be
// initialized with an empty league.
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file, %v", err)
	}

	league, err := leagueLib.NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("problem seeking start of file %s, %v", file.Name(), err)
	}

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return fmt.Errorf(
				"problem writing initial player db contents to file %s, %v",
				file.Name(),
				err,
			)
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return fmt.Errorf("problem seeking start of file %s, %v", file.Name(), err)
		}
	}

	return nil
}

// GetLeague returns the league sorted by wins in descending order.
func (f *FileSystemPlayerStore) GetLeague() leagueLib.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

// GetPlayerScore returns the score for the given player name.
// If the player does not exist, it returns 0.
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

// RecordWin records a win for the given player name.
// If the player does not exist, it creates a new player with 1 win.
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, playerLib.Player{Name: name, Wins: 1})
	}

	_ = f.database.Encode(f.league)
}
