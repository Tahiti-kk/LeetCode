package main

/**
 * 动态规划 o(n^2)
 * dp数组含义：字符串s从i到j的索引子字符串是否是回文字符串
 * 转移方程：dp[i][j] = dp[i+1][j-1] && Si==Sj
 * 转移路径：j遍历1-len(s) i遍历0-j
 */
 func longestPalindrome(s string) string {
	 length := len(s)

	 if length <= 1 {
		 return s
	 }
	 dp := make([][]bool,length)

	 // 记录最长回文子串的起始位置
	 start := 0
	 // 记录最长回文子串的最长长度
	 maxLen := 1
	 
	 // j开始向后移动
	 for j := 0; j < length; j++ {
		 dp[j] = make([]bool, length)
		 dp[j][j] = true
		 // i从0开始移动到j-1
		 for i := 0; i < j; i++ {
			 // 如果i和j的字符相等
			 if s[i] == s[j] {
				 // 特殊情况
				 // j-i = 1的情况：aa
				 // j-i = 2的情况：aba
				 if (j-i <= 2) {
					 dp[i][j] = true
				 } else {
					 // 外圈字符串与内圈字符串状态相等
					 dp[i][j] = dp[i+1][j-1]
				 }
			 } else {
				 dp[i][j] = false
			 }

			 // 如果目前子串是回文的话，与前面的状态对比
			 // 如果比前面子串长，则代替他
			 if dp[i][j] {
				curLen := j-i+1
				if curLen > maxLen {
					maxLen = curLen
					start = i
				}
			}
		 }
	 }
	 return s[start:start+maxLen]
}

func main() {
	println(longestPalindrome("abcba"))
}