package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumTotal(triangle [][]int) int {
	var n = len(triangle)
	var dp = make([][]int, n)
	for index := range dp {
		dp[index] = make([]int, index+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == n-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}
