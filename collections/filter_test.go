package collections_test

import (
	"strings"
	"testing"

	"github.com/ivnku/gobag/collections"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		collection1 := []int{1, 2, 3, 4, 5}
		collection2 := []string{"Jack", "Zack", "Bobby"}

		cond1 := func(item int) bool {
			return item%2 != 0
		}

		cond2 := func(item string) bool {
			return strings.HasSuffix(item, "ack")
		}

		t.Run("Int collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []int{1, 3, 5}, collections.Filter(collection1, cond1))
		})

		t.Run("String collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []string{"Jack", "Zack"}, collections.Filter(collection2, cond2))
		})

		t.Run("Empty collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []string{}, collections.Filter([]string{}, cond2))
		})

		t.Run("Nil collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, []string{}, collections.Filter(nil, cond2))
		})
	})
}
