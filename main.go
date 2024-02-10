package main

import (
	"data-structures-algorithms/selection_sort"
	"fmt"
)

// func RemoveIndex(s []int, index int) []int {
// 	return append(s[:index], s[index+1:]...)
// }

func main() {
	all := []int{0, 5, 8, 6, 7, 1, 2, 3, 4, 9}

	fmt.Println("all: ", all)                         //[0 1 2 3 4 6 7 8 9 9]
	fmt.Println("sorted: ", selection_sort.Sort(all)) //[0 1 2 3 4 6 7 8 9]

}

// import (
// 	binarysearch "data-structures-algorithms/binary_search"
// 	"fmt"
// )

// func main() {

// 	// queue := priorityqueue.NewPriorityQueue[string]()
// 	// queue.Enqueue("B")

// 	// queue.Enqueue("D")
// 	// queue.Enqueue("C")
// 	// queue.Enqueue("A")
// 	// fmt.Println(queue.Dequeue())
// 	// fmt.Println(queue.Dequeue())
// 	// fmt.Println(queue.Dequeue())
// 	// fmt.Println(queue.Dequeue())
// 	// fmt.Println(queue.Dequeue())

// 	data := []int{0, 1, 2, 3, 4, 5, 6, 100, 344, 6547, 999999}

// 	fmt.Println(binarysearch.Search(data, 100))
// }
