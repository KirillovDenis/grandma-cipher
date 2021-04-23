package generator

import (
	"fmt"
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

func TestWordsWeight(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		weight int
	}{
		{words: []string{"freed", "freed", "freed", "freed"}, weight: 15},
	} {
		t.Run(fmt.Sprintf("%v", tc.words), func(t *testing.T) {
			actual := computeWeights(tc.words)
			assert.Equal(t, tc.weight, actual)
		})
	}
}

func TestComputeDistance(t *testing.T) {
	for _, tc := range []struct {
		first    string
		second   string
		distance int
	}{
		{first: "freed", second: "deed", distance: 0},
		{first: "loop", second: "act", distance: 10},
	} {
		t.Run(fmt.Sprintf("distance %s->%s", tc.first, tc.second), func(t *testing.T) {
			actual := distance(tc.first, tc.second)
			assert.Equal(t, tc.distance, actual)
		})
	}
}
