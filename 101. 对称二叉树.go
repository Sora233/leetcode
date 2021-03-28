package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetric(root *TreeNode) bool {
	var nodes []*TreeNode
	nodes = append(nodes, root)
	for {
		var newNodes []*TreeNode
		for _, node := range nodes {
			if node == nil {
				continue
			}
			newNodes = append(newNodes, node.Left, node.Right)
		}
		if len(newNodes) == 0 {
			break
		}
		var f = 0
		var e = len(newNodes) - 1
		for f < e {
			if newNodes[f] == nil && newNodes[e] == nil {

			} else if newNodes[f] != nil && newNodes[e] != nil {
				if newNodes[f].Val != newNodes[e].Val {
					return false
				}
			} else {
				return false
			}
			f++
			e--
		}
		nodes = newNodes
	}
	return true
}
