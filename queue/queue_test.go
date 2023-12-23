package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue[string]()

	if queue.Length() != 0 {
		t.Error("New queue should be empty")
	}
}

func TestEnqueu(t *testing.T) {
	queue := NewQueue[string]()
	queue.Enqueue("A")

	if queue.Length() != 1 {
		t.Error("Queue should contain one value")
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueue[string]()
	queue.Enqueue("A")
	value, _ := queue.Dequeue()

	if value != "A" {
		t.Error("Dequeued value should be A")
	}
}

func TestDequeueEmpty(t *testing.T) {
	queue := NewQueue[string]()

	_, err := queue.Dequeue()

	if err == nil {
		t.Error("Dequeueing empty queue should return error")
	}
}
