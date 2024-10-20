package queue_test

import (
	"testing"

	"github.com/kiyo5hi/go-lib/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := queue.New(1, 2, 3)

	v, ok := q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	q.Enqueue(4)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, v)

	q.Enqueue(5)
	q.Enqueue(6)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 4, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 5, v)

	v, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 6, v)

	v, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Zero(t, v)
}
