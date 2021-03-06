package problem0010

func isMatch(s, p string) bool {
	sSize := len(s)
	pSize := len(p)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] 的含义是 s[:i] 可以与 p[:j] 匹配 */

	dp[0][0] = true
	/**
	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
	 * 所以，需要把相应的 dp 设置成 true
	 */
	for i := 1; i < pSize && dp[0][i-1]; i += 2 {
		if p[i] == '*' {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if p[j] == '*' {
				if p[j-1] != s[i] && p[j-1] != '.' {
					// p[j-1;j+1] 作为 "x*" 与 s[i-1;i+1] 没有匹配上
					// p[j-1:j+1] 只能被当做 ""
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					// p[j-1;j+1] 作为 "x*" 与 s[i-1;i+1] 匹配上了
					// p[j-1:j+1] 可以有三种解释
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" 解释为 "" */
						dp[i+1][j] || /* "x*" 解释为 "x" */
						dp[i][j+1] /* "x*" 解释为 "xx" */
				}
			} else if p[j] == '.' || p[j] == s[i] {
				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
				dp[i+1][j+1] = dp[i][j]
			}
		}
	}

	return dp[sSize][pSize]
}
