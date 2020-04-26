/**
 * 贪心算法
 * 每次取当前扫描的最大子序
*/
func maxSubArray(nums []int) int {
	length := len(nums)

	// curSum记录扫描到当前位置的子序和
	// maxSum记录扫描过程中的最大值
	curSum := nums[0]
    maxSum := nums[0]

	for i := 1; i < length; i++ {
		// 在 当前数 和 加上当前数的子序 中取较大的那一个
		curSum = max(nums[i], curSum+nums[i])
		// 记录最大值
		maxSum = max(curSum, maxSum)
	}
    return maxSum
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}