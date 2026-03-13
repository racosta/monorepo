package main

import "sync"

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
