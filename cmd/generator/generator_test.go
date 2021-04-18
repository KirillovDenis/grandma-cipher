package generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratorError(t *testing.T) {
	dict := []string{"a", "b"}

	generator := NewGenerator(dict)

	_, err := generator.Greedy(3, 0, 10)
	assert.Errorf(t, err, "must be too small dict error")

	_, err = generator.GreedyMult(3, 0, 10, 1)
	assert.Errorf(t, err, "must be too small dict error")
}

func TestGreedyGenerator(t *testing.T) {
	dict := []string{"hello", "lollipop", "error"}
	fmt.Println(dict)
	generator := NewGenerator(dict)

	for _, tc := range []struct {
		size     int
		min, max int
		weight   int
	}{
		{size: 1, min: 0, max: 5, weight: 11},
		{size: 1, min: 0, max: 10, weight: 8},
		{size: 2, min: 8, max: 11, weight: 26},
	} {
		t.Run(fmt.Sprintf("size: %d, min: %d, max: %d", tc.size, tc.min, tc.max), func(t *testing.T) {
			result, err := generator.Greedy(tc.size, tc.min, tc.max)
			assert.NoError(t, err)
			fmt.Println(result.Words)
			assert.Equal(t, tc.weight, result.Weight)
		})
	}

}
