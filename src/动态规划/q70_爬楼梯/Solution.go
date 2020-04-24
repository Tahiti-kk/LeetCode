func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	default:
		dp1, dp2 := 1, 2
		// 只需要记住前两次的结果
		for i := 2; i < n; i++ {
			dp1, dp2 = dp2, dp1 + dp2
		}
		return dp2
	}
}