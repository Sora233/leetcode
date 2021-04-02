package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var root = new(TreeNode)
	var mid = len(nums) / 2
	root.Val = nums[mid]
	if mid != 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid+1 != len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
