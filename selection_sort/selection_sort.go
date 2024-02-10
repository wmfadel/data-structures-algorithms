package selection_sort

func largest(nums []int) (index int) {
	index = 0
	max := nums[index]

	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	return
}

func Sort(nums []int) []int {
	sorted := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		max := largest(nums)
		sorted[i] = max
		if len(nums) > 1 {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return sorted
}
