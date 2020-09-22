package com.xiaohui.algorithm.common;

import com.xiaohui.algorithm.pojo.ListNode;

public class Utils {

    // 打印排序值
    public static String getNodeValues(ListNode startNode){
        StringBuilder result = new StringBuilder();
        ListNode currentNode = startNode;
        if(null != currentNode){
            while (currentNode.next != null){
                result.append(currentNode.val+"->");
                currentNode = currentNode.next;
            }
            result.append(currentNode.val);
            return result.toString();
        }else {
            return null;
        }
    }
}
