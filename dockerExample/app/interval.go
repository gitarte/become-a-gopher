package main

import (
	"time"
)

func interval(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
