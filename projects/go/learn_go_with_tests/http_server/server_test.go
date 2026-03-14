package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []player.Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []player.Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
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
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"

		request := testutils.NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		testutils.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []player.Player{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
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
