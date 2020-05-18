/**
 * 思路：在原有数组的前k个位置构建最小堆，最后返回nums[0]
 * 技巧：buildHeap()从倒数第一个非叶子节点开始往上堆化
 * 注意：heapify()传入的i必须为子堆堆顶
 */
func findKthLargest(nums []int, k int) int {
	buildHeap(nums, k)
	for i := k; i < len(nums); i++ {
		if nums[i] < nums[0] {
			continue
		}
		swap(nums, 0, i)
		heapify(nums, k, 0)
	}
	return nums[0]
}

func buildHeap(nums []int, k int) {
	// 从倒数第一个非叶子节点开始往上堆化
	for i:= k/2-1; i >= 0; i-- {
		heapify(nums, k, i)
	}
}

/**
 * 堆化（注意：i必须为子堆堆顶）
 * 当建堆时，前k个为规则的最小堆，通过i将堆向后扩展i
 * 当插入时，i=0，重新堆化
 */
func heapify(nums []int, k int, i int) {
	// 记录最小值的位置
	minPos := i
	for {
		minPos = min(nums, k, i, 2*i+1, 2*i+2)
		// 如果父节点已经是最小的，则退出循环，否则交换位置后扩展至下一层
		if minPos == i {
			break
		} else {
            // 交换位置
            swap(nums, i, minPos)
    		// 扩展至下一层
	    	i = minPos
        }
	}
}

/**
 * 交换数组中的两个元素
 */
func swap(nums []int, x int, y int) {
	temp := nums[x]
	nums[x] = nums[y]
	nums[y] = temp
}

// 求三个元素中最小数的位置
func min(nums []int, k int, x int, y int, z int) int {
    if y >= k {
        return x
    }
    if z >= k {
        if nums[x] <= nums[y] {
            return x
        } else {
            return y
        }
    }
    if nums[x] <= nums[y] {
        if nums[x] <= nums[z] {
            return x
        } else {
            return z
        }
    } else {
        if nums[y] <= nums[z] {
            return y
        } else {
            return z
        }
    }
}
