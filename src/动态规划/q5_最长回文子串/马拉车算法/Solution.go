package main

/**
 * 马拉车算法 o(n)
 * 中心思想：利用已知的最长回文子串直接得出
 * 			以子串内字符为中心点的子串长度
 * 辅助数组p[]，记录子串的半径
 * 数组参数：id为中心位置，mx为右边界
 * 全局参数：index为newStr的中心位置索引，maxLen为newStr中的最长半径
 * 则：
 * 子串长度：p[i]-1
 * 起始位置：(index-maxLen)/2
 * 
 */

func longestPalindrome(s string) string {

	length := len(s)

	if length <= 1 {
        return s
    }

	// 插入#
	// 开头加上$,结尾加上@
	// 目的：统一奇偶回文计算开始位置的算法
	newStr := make([]byte, 2*length + 3)
	newStr[0] = '$'
	newStr[1] = '#'
	newStr[2*length+2] = '@'
	for i:= 0; i < length; i++ {
		newStr[2*i+2] = s[i]
		newStr[2*i+3] = '#'
	}

	// 定义辅助数组p[],记录回文的半径
	p := make([]int, 2*length + 3)

	// 辅助数组中的参数
	// 回文子串中心位置
	id := 0
	// 回文子串右边界
	mx := 0

	// 全局回文参数
	// newStr中最长回文子串的中心位置索引
	index := 0
	// newStr中最长回文子串的半径
	maxLen := -1

	for i := 1; i < 2*length+1; i++ {
		// 如果i在最右边界里
		if i < mx {
			// i和2*i-id 以 id
			// 核心：p[i]要取较小的那个半径
			p[i] = min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}
		// 向两边扩展
		for newStr[i-p[i]] == newStr[i+p[i]] {
			p[i] ++
		}
		// 如果该边界超过了最右边界，则更新
		if mx < p[i] + i -1 {
			id = i
			mx = p[i] + i -1
		}
		// 如果该字串超过了目前的最长子串，则更新
		if p[i] > maxLen {
			index = i
			maxLen = p[i]
		}
	}

	start := (index-maxLen)/2
	
	return s[start:start+maxLen-1]
}


func min(x int, y int) int {
	if x<y {
        return x
    }
    return y
}

func main() {
   println(longestPalindrome("abbac"))
}