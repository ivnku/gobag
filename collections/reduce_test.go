package collections_test

import (
	"testing"

	"github.com/ivnku/gobag/collections"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {

		type Payout struct {
			ID     uint64
			Amount int
		}

		collection := []Payout{
			{
				ID:     1,
				Amount: 26,
			}, {
				ID:     2,
				Amount: 44,
			}, {
				ID:     3,
				Amount: 14,
			},
		}

		fn := func(acc int, p Payout) int {
			return acc + p.Amount
		}

		t.Run("Success reducing", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, 84, collections.Reduce(collection, 0, fn))
		})

		t.Run("Empty collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, 0, collections.Reduce([]Payout{}, 0, fn))
		})

		t.Run("Nil collection", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, 0, collections.Reduce(nil, 0, fn))
		})
	})
}
