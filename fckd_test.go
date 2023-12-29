package fckd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFckdEmptyMsg(t *testing.T) {
	assert.PanicsWithValue(t, "We are completely fckd (https://www.youtube.com/watch?v=1eknJuh8d30)", func() {
		WeAreCompletelyFckd()
	})
}

func TestFckdNonEmptyMsg(t *testing.T) {
	assert.PanicsWithValue(t, "We are completely fckd: a, b, c (https://www.youtube.com/watch?v=1eknJuh8d30)", func() {
		WeAreCompletelyFckd("a", "b", "c")
	})
}
