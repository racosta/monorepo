package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), testutils.NewPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, testutils.NewGetScoreRequest(player))
	testutils.AssertStatus(t, response.Code, http.StatusOK)

	testutils.AssertResponseBody(t, response.Body.String(), "3")
}
