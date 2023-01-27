package collections_test

import (
	"testing"

	"github.com/ivnku/gobag/collections"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {

		type Person struct {
			Name string
			Age  int
		}

		collection := []Person{
			{
				Name: "Jack",
				Age:  26,
			}, {
				Name: "Dan",
				Age:  31,
			}, {
				Name: "David",
				Age:  26,
			},
		}

		keyFn := func(p Person) int {
			return p.Age
		}

		valFn := func(p Person) Person {
			return p
		}

		expected := map[int][]Person{
			26: {
				{
					Name: "Jack",
					Age:  26,
				}, {
					Name: "David",
					Age:  26,
				},
			},
			31: {
				{
					Name: "Dan",
					Age:  31,
				},
			},
		}

		expectedEmpty := map[int][]Person{}

		t.Run("Success grouping", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, expected, collections.GroupBy(collection, keyFn, valFn))
		})

		t.Run("Empty collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, expectedEmpty, collections.GroupBy([]Person{}, keyFn, valFn))
		})

		t.Run("Nil collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, expectedEmpty, collections.GroupBy(nil, keyFn, valFn))
		})
	})
}
