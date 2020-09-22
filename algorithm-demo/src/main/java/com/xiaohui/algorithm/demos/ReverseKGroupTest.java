package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.common.Utils;
import com.xiaohui.algorithm.pojo.ListNode;

/**
 * 将给出的链表中的节点每 k 个一组翻转，返回翻转后的链表
 * 如果链表中的节点数不是 k 的倍数，将最后剩下的节点保持原样
 * 你不能更改节点中的值，只能更改节点本身。
 * 要求空间复杂度 O(1)
 * 例如：
 * 给定的链表是1->2->3->4->5
 * 对于 k=2, 你应该返回 2-> 1-> 4-> 3-> 5
 * 对于 k=3, 你应该返回 3-> 2 ->1 -> 4-> 5
 */
public class ReverseKGroupTest {
    public static void main(String[] args) {

        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        ListNode node4 = new ListNode(4);
        ListNode node5 = new ListNode(5);
        ListNode node6 = new ListNode(6);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;
        node5.next = node6;

        System.out.println("before reverse:"+ Utils.getNodeValues(node1));

        ListNode listNode = reverseKGroup(node1, 3);

        System.out.println("after reverse:"+Utils.getNodeValues(listNode));
    }

    public static ListNode reverseKGroup (ListNode head, int k) {
        if(head==null)
            return null;
        ListNode temp=head;
        int count=0;
        while(temp!=null&&count<k){
            count++;
            // 关键点1： 注意这里的temp，是第n个k组的下一个节点
            temp=temp.next;
        }
        if(count<k){
            return head;
        }
        // 关键点2：以下是在k的范围里面，将节点进行反转。
        ListNode cur=head,pre=null;
        ListNode next ;
        while(cur!=temp){
            next=cur.next;
            cur.next=pre;
            pre=cur;
            cur=next;
        }
        // 关键点3：递归，表示反转的起始节点的下一个节点是剩余节点的反转结果。
        head.next=reverseKGroup(temp,k);
        return pre;
    }

    /**
     * 查询第index位置的节点
     * @param index 从1开始
     * @param node
     * @return
     */
    public static ListNode getNode(int index, ListNode node){
        int count =0 ;
        while (node != null){
            count++;
            if(index == count){
                return node;
            }else {
                node = node.next;
            }
        }
        return null;
    }
}
