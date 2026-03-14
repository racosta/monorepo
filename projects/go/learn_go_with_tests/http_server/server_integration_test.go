package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	playerLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := testutils.CreateTempFile(t, "")
	defer cleanDatabase()
	store := NewFileSystemPlayerStore(database)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, testutils.NewGetScoreRequest(player))
		testutils.AssertStatus(t, response.Code, http.StatusOK)

		testutils.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, testutils.NewLeagueRequest())
		testutils.AssertStatus(t, response.Code, http.StatusOK)

		got := testutils.GetLeagueFromResponse(t, response.Body)
		want := []playerLib.Player{
			{Name: "Pepper", Wins: 3},
		}
		testutils.AssertLeague(t, got, want)
	})
}
