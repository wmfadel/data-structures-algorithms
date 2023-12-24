package main

import (
	priorityqueue "data-structures-algorithms/priority_queue"
	"fmt"
)

func main() {

	queue := priorityqueue.NewPriorityQueue[string]()
	queue.Enqueue("B")

	queue.Enqueue("D")
	queue.Enqueue("C")
	queue.Enqueue("A")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
