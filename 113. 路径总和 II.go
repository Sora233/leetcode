package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result [][]int
var target int

func dup(x []int) []int {
	var a []int
	for _, d := range x {
		a = append(a, d)
	}
	return a
}

func dfs(root *TreeNode, cur int, path []int) {
	cur += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if cur == target {
			result = append(result, dup(path))
			return
		}
	}
	if root.Left != nil {
		dfs(root.Left, cur, dup(path))
	}
	if root.Right != nil {
		dfs(root.Right, cur, dup(path))
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return result
	}
	result = make([][]int, 0)
	target = targetSum
	dfs(root, 0, []int{})
	return result
}
