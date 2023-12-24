package priorityqueue

import (
	"errors"
	"sync"

	"golang.org/x/exp/constraints"
)

type queue[T constraints.Ordered] struct {
	lock   sync.Mutex
	values []T
}

func NewPriorityQueue[T constraints.Ordered]() *queue[T] {
	return &queue[T]{
		lock:   sync.Mutex{},
		values: make([]T, 0),
	}
}

func (q *queue[T]) secureLock() {
	q.lock.Lock()
	defer q.lock.Unlock()
}

func (q *queue[T]) Enqueue(value T) {
	q.secureLock()

	length := len(q.values)
	if length == 0 {
		q.values = append(q.values, value)
		return
	}
	split := length
	for i, v := range q.values {
		if v > value {
			split = i
			break
		}
	}

	q.values = append(q.values[:split], append([]T{value}, q.values[split:]...)...)

}

func (q *queue[T]) Dequeue() (T, error) {
	q.secureLock()
	length := len(q.values)

	if length == 0 {
		var zeroValue T
		return zeroValue, errors.New("empty queue")
	}

	value := q.values[0]
	q.values = q.values[1:]

	return value, nil
}

func (q *queue[T]) Length() int {
	q.secureLock()
	return len(q.values)

}
