package poker

import (
	"fmt"
	"io"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	err = t.file.Truncate(0)
	if err != nil {
		return 0, fmt.Errorf("problem truncating file %s, %v", t.file.Name(), err)
	}
	ret, err := t.file.Seek(0, io.SeekStart)
	if err != nil {
		return int(ret), fmt.Errorf("problem seeking start of file %s, %v", t.file.Name(), err)
	}
	return t.file.Write(p)
}
