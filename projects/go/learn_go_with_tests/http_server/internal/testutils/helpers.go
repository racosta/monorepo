// Package testutils provides helper functions for testing the HTTP server.
package testutils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

// NewGetScoreRequest creates a new GET request for the player's score.
func NewGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

// NewPostWinRequest creates a new POST request to record a win for the player.
func NewPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

// NewLeagueRequest creates a new GET request for the league table.
func NewLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

// GetLeagueFromResponse parses the league table from the HTTP response body.
func GetLeagueFromResponse(t testing.TB, body io.Reader) (league []player.Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

// AssertStatus checks if the HTTP status code is as expected.
func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertResponseBody checks if the response body is as expected.
func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

// AssertContentType checks if the response has the expected content type.
func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

// AssertLeague checks if the league table is as expected.
func AssertLeague(t testing.TB, got, want []player.Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
