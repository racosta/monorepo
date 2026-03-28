package poker

import (
	"io"
	"testing"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/http_server/internal/testutils"
)

func TestTapeWrite(t *testing.T) {
	file, clean := testutils.CreateTempFile(t, "12345")
	defer clean()

	tape := &tape{file}

	_, err := tape.Write([]byte("abc"))
	testutils.AssertNoError(t, err)

	_, err = file.Seek(0, io.SeekStart)
	testutils.AssertNoError(t, err)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
