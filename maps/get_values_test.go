package maps_test

import (
	"testing"

	"github.com/ivnku/gobag/maps"
	"github.com/stretchr/testify/assert"
)

func TestGetValues(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			name     string
			in       map[string]int
			expected []int
		}{
			{
				name:     "nil map",
				in:       nil,
				expected: nil,
			}, {
				name:     "empty map",
				in:       map[string]int{},
				expected: []int{},
			}, {
				name: "filled map",
				in: map[string]int{
					"someKey":    14,
					"anotherKey": 66,
					"oneMoreKey": 107,
				},
				expected: []int{14, 66, 107},
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				assert.ElementsMatch(t, tc.expected, maps.GetValues(tc.in))
			})
		}
	})
}
