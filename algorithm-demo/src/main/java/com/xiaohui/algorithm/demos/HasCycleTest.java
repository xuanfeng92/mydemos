package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.common.Utils;
import com.xiaohui.algorithm.pojo.ListNode;

/**
 * 判断给定的链表中是否有环
 */
public class HasCycleTest {
    public static void main(String[] args) {
        ListNode node1 = new ListNode(111);
        ListNode node2 = new ListNode(222);
        ListNode node3 = new ListNode(333);
        ListNode node4 = new ListNode(444);
        ListNode node5 = new ListNode(555);
        ListNode node6 = new ListNode(666);
        ListNode node7 = new ListNode(777);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;
        node5.next = node6;
        node6.next = node7;
        node7.next = node3; // 这里有个环

        System.out.println("是否有环："+ hasCycle(node1));
        // System.out.println("链表："+ Utils.getNodeValues(node1)); 有环的时候，会死循环。
    }

    public static boolean hasCycle(ListNode head) {
        /*
        链表有环思路：如果有环，设置一个快指针，设置一个慢指针，
        快指针一次走两步，慢指针一次走一步，快指针总能追上慢的
        */
        if(head == null) return false;
        ListNode fast = head;
        ListNode slow = head;
        while(fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;
            if(fast == slow) {
                return true;
            }
        }
        return false;
    }
}
