// Package clockface provides functions to calculate the position of the hands of an analogue clock.
package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondsInRadians returns the angle of the second hand in radians, measured clockwise from 12 o'clock.
func SecondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

// SecondHandPoint returns the point at which the second hand of an analogue clock would end, given a time.
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// MinutesInRadians returns the angle of the minute hand in radians, measured clockwise from 12 o'clock.
func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// MinuteHandPoint returns the point at which the minute hand of an analogue clock would end, given a time.
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// HoursInRadians returns the angle of the hour hand in radians, measured clockwise from 12 o'clock.
func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

// HourHandPoint returns the point at which the hour hand of an analogue clock would end, given a time.
func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
