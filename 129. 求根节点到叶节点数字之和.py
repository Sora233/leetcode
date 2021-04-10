# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def dfs(self, root, s):
        c = s+str(root.val)
        if root.left is None and root.right is None:
            self.ans += int(c)
            return
        if root.left is not None:
            self.dfs(root.left, c)
        if root.right is not None:
            self.dfs(root.right, c)

    def sumNumbers(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        self.ans = 0
        self.dfs(root, "")
        return self.ans