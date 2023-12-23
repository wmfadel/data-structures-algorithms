package queue

import (
	"errors"
	"sync"
)

type queue struct {
	lock   sync.Mutex
	values []string
}

func NewQueue() *queue {
	return &queue{
		lock:   sync.Mutex{},
		values: make([]string, 0),
	}
}

func (q *queue) secureLock() {
	q.lock.Lock()
	defer q.lock.Unlock()
}

func (q *queue) Enqueue(value string) {
	q.secureLock()
	q.values = append(q.values, value)

}

func (q *queue) Dequeue() (string, error) {
	q.secureLock()
	length := len(q.values)

	if length == 0 {
		return "", errors.New("empty queue")
	}

	value := q.values[0]
	q.values = q.values[1:]

	return value, nil
}

func (q *queue) Length() int {
	q.secureLock()
	return len(q.values)

}
