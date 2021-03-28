package leetcode

func dfs(root *TreeNode, level int) int {
	if root == nil {
		return level - 1
	}
	var l = dfs(root.Left, level+1)
	var r = dfs(root.Right, level+1)
	var ans = l
	if r > ans {
		ans = r
	}
	return ans
}

func maxDepth(root *TreeNode) int {
	return dfs(root, 1)
}
