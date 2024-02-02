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
