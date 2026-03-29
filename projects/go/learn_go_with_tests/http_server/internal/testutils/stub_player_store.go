package testutils

import leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"

// StubPlayerStore is a test double for PlayerStore that allows us to control the data returned and track interactions.
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   leagueLib.League
}

// GetPlayerScore returns the score for the given player name from the Scores map.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin appends the player's name to the WinCalls slice to track that a win was recorded for that player.
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague returns the League field, which can be set to control the league data returned in tests.
func (s *StubPlayerStore) GetLeague() leagueLib.League {
	return s.League
}
