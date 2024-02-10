package quick_sort

func Sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]
	less := make([]int, 0)
	higher := make([]int, 0)

	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			higher = append(higher, nums[i])
		}
	}

	return append(Sort(less), append([]int{pivot}, Sort(higher)...)...)
}
