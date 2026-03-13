// Package testutils provides helper functions for testing the HTTP server.
package testutils

import (
	"fmt"
	"net/http"
	"testing"
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
