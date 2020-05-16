/**
 * 双指针题型的特点：循环两个指针，一定要考虑到边界情况，最好使用循环，不要用if语句
 */
func minSubArrayLen(s int, nums []int) int {
	// 定义左指针和右指针
	lp := 0
	result := 0
	sum := 0
	// 遍历直至rp到达终点
	for rp, value := range nums {
		sum += value
		// 如果总和大于s
		for sum >= s {
			if result == 0 {
				result = rp-lp+1
			} else{
				result = min(rp-lp+1, result)
			}
			lp++
			sum -= nums[lp-1]
		}
	}
	return result
}

func min(x int, y int) int {
	if x<y {
        return x
    }
    return y
}