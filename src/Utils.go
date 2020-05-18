func min(x int, y int) int {
	if x<y {
        return x
    }
    return y
}

func max(x int, y int) int {
	if x>y {
        return x
    }
    return y
}

/**
 * 交换数组中的两个元素
 */
 func swap(nums []int, x int, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}