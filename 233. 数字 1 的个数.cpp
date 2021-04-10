int dp[15][15];
int num[15];


int dfs(int pos, int st, bool f) {
	if (pos < 1) {
		return st;
	}
	if (!f && dp[pos][st] != -1) {
		return dp[pos][st];
	}
	int end = f ? num[pos] : 9;
	int res = 0;
	for (int i = 0; i <= end; i++) {
		bool nf = (f && i == end);
		int nst = i == 1 ? st+1 : st;
		res += dfs(pos-1, nst, nf);
	}
	if (!f) {
		dp[pos][st] = res;
	}
	return res;
}

int solve(int n) {
	int len = 0;
	while(n) {
		num[++len] = n%10;
		n /= 10;
	}
	return dfs(len, 0, 1);
}

class Solution {
public:
    int countDigitOne(int n) {
        memset(dp, -1, sizeof(dp));
		return solve(n);
    }
};