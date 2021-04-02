package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root = new(TreeNode)
	root.Val = preorder[0]

	for index, v := range inorder {
		if v == root.Val {
			if index != 0 {
				root.Left = buildTree(preorder[1:1+index], inorder[:index])
			}
			if index != len(inorder) {
				root.Right = buildTree(preorder[1+index:], inorder[index+1:])
			}
		}
	}
	return root
}
