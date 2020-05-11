package main
/**
 * 双指针
 * 遍历左指针的向右移动
 * 期间，右指针一直向右移动，直到Set中出现相同字符
 **/
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	// 左指针与右指针
	lp, rp := 0, -1
	// 记录最长长度
	maxLen := 0
	// map：key为字符，value为位置
	m := map[byte]int{}

	// 左指针向右遍历
	for ; lp <= length-1; lp++ {
		// 左指针向右移动,删除map中的字符
		if lp != 0 {
			delete(m, s[lp-1])
		}
		// 如果右一位不在窗口中，则右指针持续向右移动
		for rp<length-1 {
			_, exist := m[s[rp+1]]
			if exist {
				break
			}
			rp++
			m[s[rp]] = rp
		}
		// 记录最大值
		maxLen = max(maxLen, rp-lp+1)
	}
	return maxLen
}

func max(x int, y int) int {
	if x>y {
        return x
    }
    return y
}

func main() {
	length := lengthOfLongestSubstring("pwwkew")
	println(length)
}