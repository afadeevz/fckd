package fckd

import (
	"fmt"
	"strings"
)

func WeAreCompletelyFckd(msg ...string) {
	const (
		wacf = "We are completely fckd"
		url  = "https://www.youtube.com/watch?v=1eknJuh8d30"
	)

	var panicMsg string
	if len(msg) > 0 {
		panicMsg = fmt.Sprintf("%s: %s (%s)", wacf, strings.Join(msg, ", "), url)
	} else {
		panicMsg = fmt.Sprintf("%s (%s)", wacf, url)
	}

	panic(panicMsg)
}
