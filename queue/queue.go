package queue

type Queue[T any] struct {
	data []T
}

func New[T any](vals ...T) *Queue[T] {
	q := &Queue[T]{
		data: []T{},
	}
	for _, val := range vals {
		q.Enqueue(val)
	}
	return q
}

func (q *Queue[T]) Enqueue(val T) {
	q.data = append(q.data, val)
}

func (q *Queue[T]) Dequeue() (t T, ok bool) {
	if len(q.data) == 0 {
		return t, false
	}

	elem := q.data[0]
	q.data = q.data[1:]
	return elem, true
}

func (q *Queue[T]) Head() (t T, ok bool) {
	if len(q.data) == 0 {
		return t, false
	}
	return q.data[0], true
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}
