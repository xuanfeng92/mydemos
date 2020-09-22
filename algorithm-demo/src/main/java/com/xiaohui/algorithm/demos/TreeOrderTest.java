package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.pojo.TreeNode;
import edu.princeton.cs.algs4.In;

import java.util.ArrayList;
import java.util.List;

/**
 * 分别按照二叉树先序，中序和后序打印所有的节点。
 */
public class TreeOrderTest {
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

//        int[][] result = threeOrders(node1);
//
//        System.out.println("result:"+result);
        List<Integer> store = new ArrayList<>(getNodeCount(node1));
        printNode(store, node1);
        System.out.println("store:"+store.toString());

    }

    static int[][] res;

    static int preOrder = 0, inOrder = 0, postOrder = 0;

    public static int[][] threeOrders (TreeNode root) {
        res = new int[3][getNodeCount(root)];

        print(root);

        return res;
    }

    public static int getNodeCount(TreeNode root){
        if(root == null){
            return 0;
        }
        return 1+getNodeCount(root.left) + getNodeCount(root.right);
    }

    public static void print(TreeNode root){
        if(root == null)
            return ;
        res[0][preOrder++] = root.val;
        if(root.left != null)
            print(root.left);
        res[1][inOrder++] = root.val;
        if(root.right != null)
            print(root.right);
        res[2][postOrder++] = root.val;
    }

    public static void printNode(List<Integer> store, TreeNode root){

        if(root == null){
            return;
        }
        printNode(store, root.left);
//        store.add(root.val); // 左中右
        printNode(store, root.right);
//        store.add(root.val); // 左右中



    }

}
