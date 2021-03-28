package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverseNode(nums []int) {
	var f = 0
	var e = len(nums) - 1
	var tmp int
	for f < e {
		tmp = nums[f]
		nums[f] = nums[e]
		nums[e] = tmp
		f++
		e--
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	for level := range result {
		if level%2 == 1 {
			reverseNode(result[level])
		}
	}
	return result
}
