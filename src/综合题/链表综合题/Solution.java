/**
 * 题一：将单向链表表示十进制整数，求两个正整数的和。如1234+56=1290，用单项链表计算。
 * 注意：链表具有一个空的头结点
 */
public class Solution {
    
    /**
     * 题一：将单向链表表示十进制整数，求两个正整数的和。如1234+56=1290，用单项链表计算。
     */
    public static ListNode addIntList(ListNode head1, ListNode head2) {
        // 先将两个链表反转
        head1 = reverseList(head1);
        head2 = reverseList(head2);
        ListNode result = new ListNode(-1, null);
        // offset表示进位
        int offset = 0;
        // 将两个链表从低位开始相加
        while(head1.getNext()!=null && head2.getNext()!=null) {
            int tempInt = head1.getNext().getVal() + head2.getNext().getVal() + offset;
            offset = 0;
            if(tempInt>=10) {
                offset = 1;
                tempInt -= 10;
            }
            // 将计算后节点加入表头
            ListNode tempNode = new ListNode(tempInt, result.getNext());
            result.setNext(tempNode);
            // 删除原链表节点
            head1.setNext(head1.getNext().getNext());
            head2.setNext(head2.getNext().getNext());
        }
        while(head1.getNext() != null) {
            int tempInt = head1.getNext().getVal() + offset;
            offset = 0;
            if(tempInt>=10) {
                offset = 1;
                tempInt -= 10;
            }
            ListNode tempNode = new ListNode(tempInt, result.getNext());
            result.setNext(tempNode);
            head1.setNext(head1.getNext().getNext());
        }
        while(head2.getNext() != null) {
            int tempInt = head2.getNext().getVal() + offset;
            offset = 0;
            if(tempInt>=10) {
                offset = 1;
                tempInt -= 10;
            }
            ListNode tempNode = new ListNode(tempInt, result.getNext());
            result.setNext(tempNode);
            head2.setNext(head2.getNext().getNext());
        }
        // 如果最高位大于10
        if(offset == 1) {
            ListNode tempNode = new ListNode(1, result.getNext());
            result.setNext(tempNode);
        }
        return result;
    }

    public static void main(String[] args) {
        // 整数9999
        ListNode node4 = new ListNode(9, null);
        ListNode node3 = new ListNode(9, node4);
        ListNode node2 = new ListNode(9, node3);
        ListNode node1 = new ListNode(9, node2);
        ListNode head1 = new ListNode(-1, node1);
        // 整数99
        ListNode node6 = new ListNode(9, null);
        ListNode node5 = new ListNode(9, node6);
        ListNode head2 = new ListNode(-1, node5);

        ListNode result = addIntList(head1, head2);
        printList(result);
    }

    /**
     * 反转链表
     * 创建一个新的链表头，将旧的链表转移至新的链表
     */
    public static ListNode reverseList(ListNode head) {
        if(head==null || head.getNext()==null) {
            return head;
        }
        // 新链表的头结点
        ListNode result = new ListNode(-1, null);
        ListNode tempNode;
        // 循环加入新链表头部
        while(head.getNext() != null) {
            tempNode = head.getNext();
            head.setNext(head.getNext().getNext());
            tempNode.setNext(result.getNext());
            result.setNext(tempNode);
        }
        return result;
    }

    /**
     * 打印链表
     */
    public static void printList(ListNode head) {
        if(head == null) {
            return;
        }
        while (head.getNext() != null) {
            System.out.println(head.getNext().getVal());
            head = head.getNext();
        }
    }
}