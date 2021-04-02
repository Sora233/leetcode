package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var nodes = []*TreeNode{root}
	for {
		var newNodes []*TreeNode
		var lresult []int
		for _, node := range nodes {
			if node == nil {
				continue
			}
			lresult = append(lresult, node.Val)
			newNodes = append(newNodes, node.Left, node.Right)
		}
		if len(newNodes) == 0 {
			break
		}
		result = append(result, lresult)
		nodes = newNodes
	}
	var l = 0
	var r = len(result) - 1
	var tmp []int
	for l < r {
		tmp = result[l]
		result[l] = result[r]
		result[r] = tmp
		l++
		r--
	}
	return result
}
