package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var head *TreeNode

func do(root *TreeNode) {
	if root == nil {
		return
	}
	do(root.Right)
	do(root.Left)
	root.Right = head
	root.Left = nil
	head = root
}
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	head = nil
	do(root)
}
