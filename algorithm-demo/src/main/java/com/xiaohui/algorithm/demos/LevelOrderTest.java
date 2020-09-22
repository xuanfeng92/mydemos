package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.pojo.TreeNode;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.function.Consumer;

/**
 * 给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
 */
public class LevelOrderTest {

    public static  ArrayList<ArrayList<Integer>> levelOrder (TreeNode root){
        // 关键是如何判定换行

        //存放最终结果集
        ArrayList<ArrayList<Integer>> result = new  ArrayList<ArrayList<Integer>>();
        if(root==null) return result;
        //需要借助一个队列结构
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        //根结点入队
        queue.offer(root);
        while(!queue.isEmpty()){
            //存放每一层的节点
            ArrayList<Integer> floor = new ArrayList<Integer>();
            //本次要出队的节点数量
            int len = queue.size();
            for(int i=0;i<len;i++){
                // 关键点1：取出上层节点
                TreeNode temp = queue.poll();
                floor.add(temp.val);
                // 关键点2： 存放下层节点
                if(temp.left!=null) queue.offer(temp.left);
                if(temp.right!=null) queue.offer(temp.right);
            }
            //保存当前层结果
            result.add(floor);
        }

        return result;
    };

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


        ArrayList<ArrayList<Integer>> arrayLists = levelOrder(node1);
        arrayLists.forEach(list -> {
            System.out.println(flattenList(list));
        });
    }

    // 把list内容打印出来
    public static String flattenList(List list){
        StringBuilder result = new StringBuilder("[");
        int startCount =0 ;
        for(Object l : list){
            if(startCount == 0){
                result.append(l);
            }else{
                result.append(","+l);
            }
            startCount ++;
        }
        result.append("]");
        return  result.toString();
    }
}
