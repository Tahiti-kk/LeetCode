/**
 * 思路：
 * 通过递归判断左右节点是否与p、q相等
 * 不相等返回nil，相等返回节点
 * 如果左右节点都不是nil，说明该root节点为最近公共祖先
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	// 对root节点的左右子节点进行判断
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	// 四种情形：
	// left为nil，right为nil：返回nil
	// left或者right有一个不为nil：返回该left或者right
	// left或者right都不为nil，返回root节点
	if leftNode == nil {
		if rightNode == nil {
			return nil
		} else {
			return rightNode
		}
	} else {
		if rightNode == nil {
			return leftNode
		} else {
			return root
		}
	}
}