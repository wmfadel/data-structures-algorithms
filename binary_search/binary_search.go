package binarysearch

// Textbook binary search
func Search(a []int, search int) (result int, steps int) {
	low := 0
	high := len(a) - 1

	for low <= high {
		steps++
		result = int((low + high) / 2)
		guess := a[result]

		if guess == search {
			return
		} else if guess > search {
			high = result - 1
		} else {
			low = result + 1
		}
	}
	return
}

// func Search(a []int, search int) (result int, steps int) {
// 	mid := len(a) / 2
// 	switch {
// 	case len(a) == 0:
// 		result = -1 // not found
// 	case a[mid] > search:
// 		result, steps = Search(a[:mid], search) // right
// 	case a[mid] < search:
// 		result, steps = Search(a[mid+1:], search) // left
// 		if result >= 0 {                          // if anything but the -1 "not found" result
// 			result += mid + 1
// 		}
// 	default: // a[mid] == search
// 		result = mid // found
// 	}
// 	steps++
// 	return
// }
