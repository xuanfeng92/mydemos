package com.xiaohui.algorithm.pojo;

public class TreeNode {
    public int val = 0;
    public TreeNode left = null;
    public TreeNode right = null;


    public TreeNode(int val){
        this.val = val;
    }

    public void addNode(TreeNode left, TreeNode right){
        this.left = left;
        this.right = right;
    }
}
