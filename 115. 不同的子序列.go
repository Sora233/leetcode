package leetcode

func numDistinct(s string, t string) int {
	var dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	var slen = len(s)
	var tlen = len(t)
	for i := 0; i <= slen; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= slen; i++ {
		for j := 1; j <= tlen; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
			dp[i][j] += dp[i-1][j]
		}
	}
	return dp[slen][tlen]
}
