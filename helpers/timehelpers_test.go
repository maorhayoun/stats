package helpers

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestCollectDateKeys(t *testing.T) {

	var dateKeysCases = []struct {
		testDate       time.Time
		hours          int
		expectedLength int
		expectedValues []interface{}
	}{
		{time.Date(2015, time.June, 19, 10, 20, 0, 0, time.UTC), 1, 1, []interface{}{"201506191000"}},
		{time.Date(2015, time.June, 19, 10, 20, 0, 0, time.UTC), 2, 2, []interface{}{"201506191000", "201506190900"}},
		{time.Date(2015, time.June, 19, 0, 0, 0, 0, time.UTC), 2, 2, []interface{}{"201506190000", "201506182300"}},
	}

	for _, tc := range dateKeysCases {
		actual := CollectDateKeys(tc.testDate, tc.hours)
		assert.Equal(t, len(actual), tc.expectedLength)
		for i := 0; i < len(actual); i++ {
			assert.Equal(t, tc.expectedValues[i], actual[i])
		}
	}
}
