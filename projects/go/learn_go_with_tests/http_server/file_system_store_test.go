package main

import (
	"testing"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader, sorted", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)

		got := store.GetLeague()

		want := leagueLib.League{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		testutils.AssertLeague(t, got, want)

		got = store.GetLeague()
		testutils.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		testutils.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := testutils.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		testutils.AssertNoError(t, err)
	})
}
