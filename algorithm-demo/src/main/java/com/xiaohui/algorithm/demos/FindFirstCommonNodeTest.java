package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.pojo.ListNode;

/**
 * 输入两个链表，找出它们的第一个公共结点。
 */
public class FindFirstCommonNodeTest {
    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        ListNode node4 = new ListNode(4);
        ListNode node5 = new ListNode(5);
        ListNode node6 = new ListNode(6);

        ListNode node11 = new ListNode(11);
        ListNode node22 = new ListNode(22);
        ListNode node33 = new ListNode(33);
        ListNode node44 = new ListNode(44);
        ListNode node55 = new ListNode(55);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;
        node5.next = node6;

        node11.next = node22;
        node22.next = node33;
        node33.next = node44;
        node44.next = node55;
        node55.next = node4;  // 这里是公共点


        assert findFirstCommonNode(node1, node11) != null;
        System.out.println("找到的公共节点："+ findFirstCommonNode(node1, node11).val);
    }

    public static ListNode findFirstCommonNode(ListNode pHead1, ListNode pHead2) {
        ListNode cur1 = pHead1;
        ListNode cur2 = pHead2;
        while(cur1 != null){
            while(cur2 != null){
                if(cur2 == cur1){
                    return cur1;
                }else{
                    cur2 = cur2.next;
                }
            }
            cur1 = cur1.next;
            cur2 = pHead2;
        }
        return null;
    }
}
