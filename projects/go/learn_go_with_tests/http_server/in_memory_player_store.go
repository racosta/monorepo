package main

import (
	"sync"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
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

func (s *InMemoryPlayerStore) GetLeague() []player.Player {
	var league []player.Player
	for name, wins := range s.store {
		league = append(league, player.Player{Name: name, Wins: wins})
	}
	return league
}
