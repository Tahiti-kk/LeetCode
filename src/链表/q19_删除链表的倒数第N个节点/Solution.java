/**
 * 使用双指针，rp前进n个元素，当rp到末尾时，lp到达要删除的元素
 * 
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        // 链表里只有一个元素时删除这个元素
        if(head.next == null) {
            head = null;
            return head;
        }
        ListNode lp = head, rp = head;
        for(int i = 0; i < n; i++) {
            rp = rp.next;
        }
        // 要删除第一个元素的情况
        if(rp == null) {
            return head.next;
        }
        while(rp.next != null) {
            lp = lp.next;
            rp = rp.next;
        }
        lp.next = lp.next.next;
        return head;
    }
}