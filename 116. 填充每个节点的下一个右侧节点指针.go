package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var nodes = []*Node{root}
	for {
		var newNodes []*Node
		for _, node := range nodes {
			if node == nil {
				continue
			}
			newNodes = append(newNodes, node.Left, node.Right)
		}
		for index := range newNodes {
			if index+1 == len(newNodes) {
				break
			}
			if newNodes[index] == nil {
				break
			}
			newNodes[index].Next = newNodes[index+1]
		}
		if len(newNodes) == 0 {
			break
		}
		nodes = newNodes
	}
	return root
}
