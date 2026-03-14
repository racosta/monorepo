package main

import (
	"encoding/json"
	"io"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   leagueLib.League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	_, _ = database.Seek(0, io.SeekStart)
	league, _ := leagueLib.NewLeague(database)
	return &FileSystemPlayerStore{
		database: database,
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() leagueLib.League {
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

	_, _ = f.database.Seek(0, io.SeekStart)
	_ = json.NewEncoder(f.database).Encode(f.league)
}
