#include <algorithm>

const int N = 3e4+5;
const int inf = ~1u>>3;
class Solution {
private:
    int ans;
    void init() {
        ans = -inf;
    }
public:
    int dfs(TreeNode* r) {
        if (r == NULL) {
            return -inf;
        }
        int a1 = dfs(r->left);
        int a2 = dfs(r->right);
        ans = max(ans, a1);
        ans = max(ans, a2);
        ans = max(ans, a1+a2 + r->val);
        ans = max(ans, a1+r->val);
        ans = max(ans, a2+r->val);
        ans = max(ans, r->val);
        return max(r->val, max(a1+r->val, a2+r->val));
    }
    int maxPathSum(TreeNode* root) {
        init();
        dfs(root);
        return ans;
    }
};