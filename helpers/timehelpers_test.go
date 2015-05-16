package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var timeCases = []struct {
	in  string
	out int
}{
	{"1d", 24},
	{"5h", 5},
	{"2d3h", 51},
	{"4", 4},
	{"", 0},
	{"asdfsdf", 0},
}

func TestParseHours(t *testing.T) {
	for _, tc := range timeCases {
		n := ParseHours(tc.in)
		assert.Equal(t, tc.out, n)
	}
}

