/**
 * 动态规划
 * 状态定义：dp[i]表示以nums[i]结尾的最大子序和
 * 状态转移：dp[i] = dp[i] = max(dp[i]+nums[i], nums[i])
 */
func maxSubArray(nums []int) int {
	length := len(nums)

	// 状态定义：dp[i]表示以nums[i]结尾的最大子序和
	dp := make([]int, length)

	// 状态初始化
	dp[0] = nums[0]

	// 状态转移
	for i:=1; i < length; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	// 取dp中最大的数
	maxSum := dp[0]
	for i:=1; i < length; i++ {
		maxSum = max(maxSum, dp[i])
	}
	return maxSum;
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}