// Package main provides a simple program to render an SVG clockface showing the current time.
package main

import (
	"os"
	"time"

	svg "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
