package priorityqueue

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	q := NewPriorityQueue[int]()

	if q.Length() != 0 {
		t.Errorf("New queue should be empty")
	}
}

func TestEnqueue(t *testing.T) {
	q := NewPriorityQueue[int]()

	q.Enqueue(1)
	q.Enqueue(21)
	q.Enqueue(4)

	if q.Length() != 3 {
		t.Errorf("Queue should contain one value")
	}

}

func TestDequeue(t *testing.T) {
	q := NewPriorityQueue[int]()

	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)

	first, _ := q.Dequeue()
	second, _ := q.Dequeue()
	third, _ := q.Dequeue()

	if first != 1 || second != 2 || third != 3 {
		t.Errorf("Dequeued values are in wrong order")
	}

}

func TestDequeueEmpty(t *testing.T) {
	q := NewPriorityQueue[int]()

	_, err := q.Dequeue()

	if err == nil {
		t.Errorf("Dequeue from empty queue should return error")
	}

}
