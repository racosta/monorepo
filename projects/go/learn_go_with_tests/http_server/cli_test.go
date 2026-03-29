package poker_test

import (
	"strings"
	"testing"

	poker "github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server"
	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &testutils.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		testutils.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &testutils.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		testutils.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
