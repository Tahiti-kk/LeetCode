/**
 * 左右各有一个指针
 * 指针每次向小的那个数移动
 */
func maxArea(height []int) int {
    length := len(height)
    // 定义左指针与右指针
    lp, rp := 0, length-1
    // 记录最大值
    maxCap := 0
    for lp < rp {
        capacity := min(height[lp], height[rp]) * (rp-lp)
        maxCap = max(maxCap, capacity)
        if height[lp] < height[rp] {
            lp++
        } else {
            rp--
        }
    }
    return maxCap
}

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