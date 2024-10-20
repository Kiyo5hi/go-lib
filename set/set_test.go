package set_test

import (
	"testing"

	"github.com/kiyo5hi/go-lib/set"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := set.New(1)
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(0))

	s.Add(1)
	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Contains(1))

	s.Remove(1)
	assert.Equal(t, 0, s.Len())
	assert.False(t, s.Contains(1))

	s1 := set.New(1, 2, 3)
	s2 := set.New(2, 3, 4)
	unioned := s1.Union(s2)
	assert.Equal(t, 4, unioned.Len())
	for _, v := range []int{1, 2, 3, 4} {
		assert.True(t, unioned.Contains(v))
	}
	assert.ElementsMatch(t, []int{1, 2, 3, 4}, unioned.ToSlice())
}
