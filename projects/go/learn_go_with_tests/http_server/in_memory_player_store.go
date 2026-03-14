package main

import (
	"sync"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[name]++
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.store[name]
}

func (s *InMemoryPlayerStore) GetLeague() leagueLib.League {
	var league leagueLib.League
	for name, wins := range s.store {
		league = append(league, playerLib.Player{Name: name, Wins: wins})
	}
	return league
}
