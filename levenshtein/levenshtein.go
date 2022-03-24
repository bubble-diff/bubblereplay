package levenshtein

// Compute 计算两返回差异百分比
func Compute(resp1, resp2 string) float64 {
	minDistance := float64(levenshtein(resp1, resp2))
	maxLen := float64(max(len(resp1), len(resp2)))
	return (minDistance / maxLen) * 100
}

// levenshtein 编辑距离算法
func levenshtein(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]-1)
			} else {
				dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
			}
		}
	}
	return dp[m][n]
}

// min 返回三值中最小值
func min(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return a
}

// max 返回两值较大者
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
