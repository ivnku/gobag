package collections_test

import (
	"testing"

	"github.com/ivnku/gobag/collections"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
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

		mapFn := func(p Person) string {
			return p.Name
		}

		expected := []string{"Jack", "Dan", "David"}

		t.Run("Success mapping", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, expected, collections.Map(collection, mapFn))
		})

		t.Run("Empty collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []string{}, collections.Map([]Person{}, mapFn))
		})

		t.Run("Nil collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []string{}, collections.Map(nil, mapFn))
		})
	})
}
