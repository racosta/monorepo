package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestGETPlayers(t *testing.T) {
	store := testutils.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		WinCalls: nil,
		League:   nil,
	}
	server := NewPlayerServer(&store)

	tests := []struct {
		name               string
		player             string
		expectedHTTPStatus int
		expectedScore      string
	}{
		{
			name:               "returns Pepper's score",
			player:             "Pepper",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "returns Floyd's score",
			player:             "Floyd",
			expectedHTTPStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "returns 404 on missing players",
			player:             "Apollo",
			expectedHTTPStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := testutils.NewGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			testutils.AssertStatus(t, response.Code, tt.expectedHTTPStatus)
			testutils.AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}
}

func TestStoreWins(t *testing.T) {
	store := testutils.StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: nil,
		League:   nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := testutils.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		testutils.AssertStatus(t, response.Code, http.StatusAccepted)

		testutils.AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := leagueLib.League{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := testutils.StubPlayerStore{Scores: nil, WinCalls: nil, League: wantedLeague}
		server := NewPlayerServer(&store)

		request := testutils.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := testutils.GetLeagueFromResponse(t, response.Body)
		testutils.AssertStatus(t, response.Code, http.StatusOK)
		testutils.AssertContentType(t, response, jsonContentType)
		testutils.AssertLeague(t, got, wantedLeague)
	})
}
