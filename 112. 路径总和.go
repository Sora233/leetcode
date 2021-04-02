package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var target int

func dfs(root *TreeNode, cur int) bool {
	cur += root.Val
	if root.Left == nil && root.Right == nil {
		return cur == target
	}
	if root.Left != nil {
		if dfs(root.Left, cur) {
			return true
		}
	}
	if root.Right != nil {
		if dfs(root.Right, cur) {
			return true
		}
	}
	return false
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	target = targetSum
	return dfs(root, 0)
}
