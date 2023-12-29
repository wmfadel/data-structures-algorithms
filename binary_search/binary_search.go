package binarysearch

import "fmt"

func Search(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = Search(a[:mid], search)
	case a[mid] < search:
		result, searchCount = Search(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
			fmt.Println("result", result)
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}
