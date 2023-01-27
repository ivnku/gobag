package maps_test

import (
	"testing"

	"github.com/ivnku/gobag/maps"
	"github.com/stretchr/testify/assert"
)

func TestReform(t *testing.T) {
	t.Parallel()

	type Person struct {
		Name string
		Age  int
	}

	type Position struct {
		ID    uint64
		Title string
	}

	t.Run("Success", func(t *testing.T) {
		testCases := []struct {
			name     string
			in       map[Person]Position
			expected map[string]uint64
			fnKey    func(Person) string
			fnVal    func(Position) uint64
		}{
			{
				name:     "nil map",
				in:       nil,
				expected: map[string]uint64{},
				fnKey: func(p Person) string {
					return p.Name
				},
				fnVal: func(pos Position) uint64 {
					return pos.ID
				},
			}, {
				name:     "empty map",
				in:       map[Person]Position{},
				expected: map[string]uint64{},
				fnKey: func(p Person) string {
					return p.Name
				},
				fnVal: func(pos Position) uint64 {
					return pos.ID
				},
			}, {
				name: "filled map",
				in: map[Person]Position{
					{
						Name: "John",
						Age:  26,
					}: {
						ID:    1444,
						Title: "Engineer",
					}, {
						Name: "Lionel",
						Age:  35,
					}: {
						ID:    1010,
						Title: "Football player",
					},
				},
				expected: map[string]uint64{
					"John":   1444,
					"Lionel": 1010,
				},
				fnKey: func(p Person) string {
					return p.Name
				},
				fnVal: func(pos Position) uint64 {
					return pos.ID
				},
			},
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()
				assert.Equal(t, tc.expected, maps.Reform(tc.in, tc.fnKey, tc.fnVal))
			})
		}
	})
}
