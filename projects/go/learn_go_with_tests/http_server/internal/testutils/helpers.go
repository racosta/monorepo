// Package testutils provides helper functions for testing the HTTP server.
package testutils

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	leagueLib "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/league"
)

// NewGetScoreRequest creates a new GET request for the player's score.
func NewGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"/players/"+name,
		nil,
	)
	return request
}

// NewPostWinRequest creates a new POST request to record a win for the player.
func NewPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/players/"+name,
		nil,
	)
	return request
}

// NewLeagueRequest creates a new GET request for the league table.
func NewLeagueRequest() *http.Request {
	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "/league", nil)
	return request
}

// GetLeagueFromResponse parses the league table from the HTTP response body.
func GetLeagueFromResponse(tb testing.TB, body io.Reader) (league leagueLib.League) {
	tb.Helper()
	league, _ = leagueLib.NewLeague(body)
	return league
}

// AssertNoError checks if err is nil as expected.
func AssertNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("didn't expect an error but got one, %v", err)
	}
}

// AssertStatus checks if the HTTP status code is as expected.
func AssertStatus(tb testing.TB, got, want int) {
	tb.Helper()
	if got != want {
		tb.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertResponseBody checks if the response body is as expected.
func AssertResponseBody(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

// AssertContentType checks if the response has the expected content type.
func AssertContentType(tb testing.TB, response *httptest.ResponseRecorder, want string) {
	tb.Helper()
	if response.Result().Header.Get("content-type") != want {
		tb.Errorf(
			"response did not have content-type of %s, got %v",
			want,
			response.Result().Header,
		)
	}
}

// AssertLeague checks if the league table is as expected.
func AssertLeague(tb testing.TB, got, want leagueLib.League) {
	tb.Helper()

	if !reflect.DeepEqual(got, want) {
		tb.Errorf("got %v want %v", got, want)
	}
}

// AssertScoreEquals checks if the player's score is as expected.
func AssertScoreEquals(tb testing.TB, got, want int) {
	tb.Helper()
	if got != want {
		tb.Errorf("got score %d want %d", got, want)
	}
}

// AssertPlayerWin checks if a win was recorded for the expected player.
func AssertPlayerWin(tb testing.TB, store *StubPlayerStore, winner string) {
	tb.Helper()

	if len(store.WinCalls) == 0 {
		tb.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
	}

	got := store.WinCalls[0]
	if got != winner {
		tb.Errorf("did not store correct winner got %q want %q", got, winner)
	}
}

// CreateTempFile creates a temporary file with provided data and a function to cleanup
func CreateTempFile(tb testing.TB, initialData string) (*os.File, func()) {
	tb.Helper()

	tmpfile, err := os.CreateTemp("", "db")

	if err != nil {
		tb.Fatalf("could not create temp file %v", err)
	}

	_, _ = tmpfile.Write([]byte(initialData))

	removeFile := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name()) //nolint:gosec
	}

	return tmpfile, removeFile
}
