package stack_test

import (
	"testing"

	"github.com/kiyo5hi/go-lib/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack.New(1, 2, 3)

	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	s.Push(4)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = s.Pop()
	assert.False(t, ok)
	assert.Zero(t, v)
}
