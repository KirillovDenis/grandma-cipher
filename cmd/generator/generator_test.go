package generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratorNoRecursive(t *testing.T) {
	for _, tc := range []struct {
		dict   []string
		weight int
	}{
		{dict: []string{"lollipop", "ssw", "freed", "weeds", "redder", "weeder", "reeds", "poop", "loop", "pool"}, weight: 12},
	} {
		t.Run("test", func(t *testing.T) {
			gen := NewGenerator(tc.dict)
			res := gen.NoRecursive(4, 20, 24, 3)
			fmt.Println(res)
			assert.Equal(t, tc.weight, res.Weight)
		})
	}
}
