package collections_test

import (
	"github.com/ivnku/gobag/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIn(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		collection1 := []int{1, 2, 3, 4, 5}
		collection2 := []string{"Jack", "Zack", "Bobby"}

		t.Run("Int collection true", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, true, collections.In(collection1, 4))
		})

		t.Run("String collection true", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, true, collections.In(collection2, "Bobby"))
		})

		t.Run("Int collection false", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, false, collections.In(collection1, 11))
		})

		t.Run("String collection false", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, false, collections.In(collection2, "Lionel"))
		})

		t.Run("Empty collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, false, collections.In([]string{}, "Cristiano"))
		})

		t.Run("Nil collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, false, collections.In(nil, "Frank"))
		})
	})
}
