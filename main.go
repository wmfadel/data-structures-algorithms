package main

import (
	"data-structures-algorithms/queue"
	"fmt"
)

func main() {

	queue := queue.NewQueue[string]()

	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Enqueue("C")
	queue.Enqueue("D")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
