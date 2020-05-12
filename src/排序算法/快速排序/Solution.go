/**
 * 第一个数为基准，左右指针相向运动，交换不符合基准的左右两个数，并进行递归
 * 坑：需要先将右指针向左移动，以避免发生最左端是最小数的情况！！!
 */
func quickSort(array []int, low int, high int) {
	if low >= high {
        return
    }
	// 基准位
	temp := array[low]
	lp, rp := low, high
	for lp<rp {
		// 如果左指针小于等于基准，则向右移动左指针
		// 如果左指针大于基准，则停止移动
		for array[rp] >= temp && lp<rp {
			rp--
		}
		for array[lp] <= temp && lp<rp {
			lp++
		}
		if lp<rp {
			// 交换两个指针指向的数
			t := array[lp]
			array[lp] = array[rp]
			array[rp] = t
		}
	}
	// 至此，lp左边的数小于等于基准，rp右边的数大于等于基准
	// 将lp指向的数与基准置换
	array[low] = array[lp]
	array[lp] = temp
	// 递归调用左右两边数组
	quickSort(array, low, lp-1)
	quickSort(array, lp+1, high)
}