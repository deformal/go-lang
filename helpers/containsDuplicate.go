package helpers

func ContainsDuplicate(nums []int) bool {
	result := QuickSort(nums, 0, len(nums)-1)
	if len(result) <= 1 {
		return false
	}
	for i := 0; i < len(result)-1; i++ {
		current := nums[i]
		next := nums[i+1]
		if current == next {
			return true
		}
	}
	return false
}
