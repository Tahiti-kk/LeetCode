/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
	 return recur(root) != -1
}

// 自底至顶
// 可以提前阻断，性能更好
// 得到一层就判断一层，及时将信息抛出
func recur(root *TreeNode) int {
	if root == nil {
		return nil
	}
	// 如果左子树不为平衡树，直接返回-1
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	// 如果右子树不为平衡树，直接返回-1
	right := recur(root.Right)
	if right == -1 {
		return -1
	}
	// 如果左右子树都为平衡树，则返回该节点的高度
	if left-right >= -1 && left-right <= 1 {
		return max(left, right)+1
	} else {
		return -1
	}

}

func max(x int, y int) int {
	if x>y {
        return x
    }
    return y
}