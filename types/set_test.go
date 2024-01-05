package types_test

import (
	"testing"

	"github.com/ivnku/gobag/types"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Parallel()

	t.Run("Set tests", func(t *testing.T) {
		ints := []int{11, 22, 33}

		setFromNil := types.NewSet[string](nil)
		setFromInts := types.NewSet(ints)

		setFromNil.Add("test")
		assert.Equal(t, true, setFromNil.IsExists("test"))
		assert.Equal(t, 1, setFromNil.Len())
		assert.Equal(t, []string{"test"}, setFromNil.ToSlice())
		assert.Equal(t, 0, setFromNil.Remove("test").Len())
		assert.Equal(t, 3, setFromNil.AddMultiple([]string{"a", "b", "c"}).Len())

		setFromInts.Add(66)
		assert.Equal(t, true, setFromInts.IsExists(11))
		assert.Equal(t, 4, setFromInts.Len())
		assert.ElementsMatch(t, []int{11, 22, 33, 66}, setFromInts.ToSlice())
		assert.Equal(t, 3, setFromInts.Remove(22).Len())
		assert.Equal(t, 6, setFromInts.AddMultiple([]int{100, 101, 102}).Len())
	})
}
