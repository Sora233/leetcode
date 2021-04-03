package leetcode

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var set = make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	nums = make([]int, 0)
	for k := range set {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	var cur = 1
	var ans = 1
	for index := range nums {
		if index == 0 {
			continue
		}
		if nums[index] == nums[index-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
