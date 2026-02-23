// Package svg provides a function to write an SVG representation of an analogue clock showing a given time.
package svg

import (
	"fmt"
	"io"
	"time"

	clockface "github.com/racosta/monorepo/projects/go/learn_go_with_tests/math"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
)

// Write writes an SVG representation of an analogue clock showing the time `t` to `w`.
func Write(w io.Writer, t time.Time) {
	_, _ = io.WriteString(w, svgStart)
	_, _ = io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	_, _ = io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.SecondHandPoint(t), secondHandLength)
	_, _ = fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.MinuteHandPoint(t), minuteHandLength)
	_, _ = fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(clockface.HourHandPoint(t), hourHandLength)
	_, _ = fmt.Fprintf(
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:5px;"/>`,
		p.X,
		p.Y,
	)
}

func makeHand(p clockface.Point, length float64) clockface.Point {
	p = clockface.Point{X: p.X * length, Y: p.Y * length}
	p = clockface.Point{X: p.X, Y: -p.Y}
	return clockface.Point{X: p.X + clockCenterX, Y: p.Y + clockCenterY}
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">
`

const bezel = `  <circle cx="150" cy="150" r="100" style="fill: none; stroke: black; stroke-width: 5px"/>
`

const svgEnd = `</svg>`
