package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead, temp *ListNode

	for head != nil {
		// 使用temp暂存head.Next
		temp = head.Next
		// 转移原链表表头至新链表
		head.Next = newHead
		newHead = head
		// 修改原链表
		head = temp
	}
	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

func main() {
	var node1, node2, node3, node4, node5 ListNode
	var head *ListNode = &node1
	node1 = ListNode{1, &node2}
	node2 = ListNode{2, &node3}
	node3 = ListNode{3, &node4}
	node4 = ListNode{4, &node5}
	node5 = ListNode{5, nil}
	printList(head)
	head = reverseList(head)
	printList(head)
}