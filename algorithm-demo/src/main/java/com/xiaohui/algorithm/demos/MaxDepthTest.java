package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.pojo.TreeNode;

import java.util.LinkedList;
import java.util.Queue;

/**
 * 求给定二叉树的最大深度，
 * 最大深度是指树的根结点到最远叶子结点的最长路径上结点的数量。
 */
public class MaxDepthTest {

    public static int maxDepth (TreeNode root) {

        // 解法一：通过遍历可用节点
        if(root == null) return 0;
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offer(root);
        int count = 0;
        while (!queue.isEmpty()){
            // 每次循环表示一个深度
            count ++;
            //本次要出队的节点数量
            int len = queue.size();
            for(int i=0;i<len;i++){
                // 关键点1：取出上层节点
                TreeNode temp = queue.poll();
                // 关键点2： 存放下层节点
                if(temp.left!=null) queue.offer(temp.left);
                if(temp.right!=null) queue.offer(temp.right);
            }
        }
        return count;

        // 解法二：通过递归
//        if(root == null)return 0;
//
//        int leftDepth = maxDepth(root.left);
//        int rightDepth = maxDepth(root.right);
//        return 1 + Math.max(leftDepth, rightDepth);
    }
    public static void main(String[] args) {
        TreeNode node1 = new TreeNode(3);
        TreeNode node2 = new TreeNode(9);
        TreeNode node3 = new TreeNode(20);
        TreeNode node4 = new TreeNode(15);
        TreeNode node5 = new TreeNode(7);
        TreeNode node6 = new TreeNode(9);
        TreeNode node7 = new TreeNode(17);
        TreeNode node8 = new TreeNode(23);
        TreeNode node9 = new TreeNode(29);

        node1.addNode(node2,node3);
        node2.addNode(node4,node5);
        node3.addNode(node6,node7);
        node4.addNode(node8,node9);

        System.out.println("最大深度："+ maxDepth(node1));


    }
}
