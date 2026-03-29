package poker

import (
	"sync"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

// InMemoryPlayerStore is an in-memory implementation of PlayerStore.
type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

// NewInMemoryPlayerStore initializes and returns a new InMemoryPlayerStore.
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

// RecordWin records a win for the given player name.
// If the player does not exist, it creates a new player with 1 win.
func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[name]++
}

// GetPlayerScore returns the score for the given player name.
func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.store[name]
}

// GetLeague returns the league sorted by wins in descending order.
func (s *InMemoryPlayerStore) GetLeague() leagueLib.League {
	var league leagueLib.League
	for name, wins := range s.store {
		league = append(league, playerLib.Player{Name: name, Wins: wins})
	}
	return league
}
