package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	isEmpty := stack.IsEmpty()

	if isEmpty == false {
		t.Errorf("New Stack should be empty")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	isEmpty := stack.IsEmpty()

	if isEmpty == true {
		t.Errorf("Value should be added to stack")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	value, err := stack.Pop()
	if err != nil {
		t.Errorf("Failed to pop value")
	}
	if value == 0 {
		t.Errorf("Value should be added to stack")
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := NewStack[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Should return error when poping an empty stack")
	}

}
func TestPeek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	value := stack.Peek()

	if value != 1 {
		t.Errorf("Value should be added to stack")
	}
}
