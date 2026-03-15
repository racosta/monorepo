package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

type FileSystemPlayerStore struct {
	database *json.Encoder
	league   leagueLib.League
}

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

func (f *FileSystemPlayerStore) GetLeague() leagueLib.League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, playerLib.Player{Name: name, Wins: 1})
	}

	_ = f.database.Encode(f.league)
}
