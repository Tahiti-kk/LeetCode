func threeSum(nums []int) [][]int {
	length := len(nums)
	if nums == nil || length<3 {
		return nil
	}
	// 先进行排序
    quickSort(nums, 0, length-1)
    var result [][]int

    // 以第一个数为基准，向后遍历
    for i:=0; nums[i]<=0 && i<length-1; i++ {
		// 如果重复，则跳过，避免重复的情况
		if i>0 && nums[i]==nums[i-1] {
			continue
		}
		lp, rp := i+1, length-1
		for lp<rp {
			// 如果三数之和小于0，则左指针右移，否则右指针左移
			if nums[i]+nums[lp]+nums[rp] < 0 {
				lp++
			} else if nums[i]+nums[lp]+nums[rp] > 0 {
				rp--
			} else {
				// 符合条件，加入到集合中
				result = append(result, []int{nums[i], nums[lp], nums[rp]})
				// 指针指向的数，如果重复，则跳过
				for lp<rp && nums[lp]==nums[lp+1] {
					lp++
				}
				for lp<rp && nums[rp]==nums[rp-1] {
					rp--
				}
				lp++
				rp--
			}
		}
	}
	return result
}

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