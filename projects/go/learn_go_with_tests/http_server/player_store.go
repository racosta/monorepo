package poker

import leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() leagueLib.League
}
