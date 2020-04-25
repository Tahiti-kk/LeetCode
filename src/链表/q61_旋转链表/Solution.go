/**
 * 旋转链表:将链表每个节点向右移动 k 个位置，其中 k 是非负数。
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

	// 将链表连成环
	// 并记录节点个数size
    var size int
    oldTail := head
    for size = 1; oldTail.Next != nil; size++ {
        oldTail = oldTail.Next
    }
    oldTail.Next = head
	
	// 循环至断点前一个节点
    for i := 0; i < size - k%size - 1; i++ {
        head = head.Next
    }
    newHead := head.Next
    head.Next = nil
    return newHead
}