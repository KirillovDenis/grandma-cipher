package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordWeight(t *testing.T) {
	for _, tc := range []struct {
		word   string
		weight int
	}{
		{word: "hello", weight: 12},
		{word: "qwerty", weight: 5},
		{word: "a", weight: 0},
		{word: "ffff", weight: 0},
		{word: "mlp", weight: 5},
		{word: "zpmq", weight: 24},
		{word: "freed", weight: 3},
		{word: "error", weight: 11},
	} {
		t.Run(tc.word, func(t *testing.T) {
			actual := weight(tc.word)
			assert.Equal(t, tc.weight, actual)
		})
	}
}
