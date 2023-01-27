package maps_test

import (
	"testing"

	"github.com/ivnku/gobag/maps"
	"github.com/stretchr/testify/assert"
)

func TestGetKeys(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			name     string
			in       map[string]int
			expected []string
		}{
			{
				name:     "nil map",
				in:       nil,
				expected: nil,
			}, {
				name:     "empty map",
				in:       map[string]int{},
				expected: []string{},
			}, {
				name: "filled map",
				in: map[string]int{
					"someKey":    14,
					"anotherKey": 66,
					"oneMoreKey": 107,
				},
				expected: []string{"someKey", "anotherKey", "oneMoreKey"},
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				assert.ElementsMatch(t, tc.expected, maps.GetKeys(tc.in))
			})
		}
	})
}
