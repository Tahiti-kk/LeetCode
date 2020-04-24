func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	if numRows == 0 {
		return result
	}
    // 一共numRows行
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, 0)
        // 一行i个元素
		for j :=0; j <= i; j++ {
			if j == 0 || j == i {
				result[i] = append(result[i], 1)
			} else {
				result[i] = append(result[i],result[i-1][j-1] + result[i-1][j])
			}
		}
	}
	return result
}