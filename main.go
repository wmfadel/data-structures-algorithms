package main

import (
	binarysearch "data-structures-algorithms/binary_search"
	"fmt"
)

func main() {

	// queue := priorityqueue.NewPriorityQueue[string]()
	// queue.Enqueue("B")

	// queue.Enqueue("D")
	// queue.Enqueue("C")
	// queue.Enqueue("A")
	// fmt.Println(queue.Dequeue())
	// fmt.Println(queue.Dequeue())
	// fmt.Println(queue.Dequeue())
	// fmt.Println(queue.Dequeue())
	// fmt.Println(queue.Dequeue())

	data := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(binarysearch.Search(data, 6))
}
