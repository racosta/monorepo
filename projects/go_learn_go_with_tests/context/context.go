// Package context implements a simple HTTP server with cancelable operations.
package context

import (
	"context"
	"fmt"
	"net/http"
)

// A Store has a long-running task to Fetch that can be cancelled.
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server creates HTTP handler functions for long-running Fetch requests.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // TODO: log error however you like
		}

		_, _ = fmt.Fprint(w, data)
	}
}
