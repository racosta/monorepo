// Package league implements the League used in the application.
package league

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/player"
)

// League represents a collection of Players
type League []player.Player

// NewLeague creates a League instance from a Reader
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("problem parsing league, %v", err)
	}
	return league, nil
}

// Find searches the League collection for a player whose name matches `name` and returns that Player instance if it exists.
func (l League) Find(name string) *player.Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
