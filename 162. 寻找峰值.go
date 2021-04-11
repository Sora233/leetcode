package leetcode

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for index := range nums {
		if index == 0 {
			if nums[index] > nums[index+1] {
				return index
			}
		} else if index+1 == len(nums) {
			if nums[index] > nums[index-1] {
				return index
			}
		} else {
			if nums[index] > nums[index-1] && nums[index] > nums[index+1] {
				return index
			}
		}
	}
	return 0
}
