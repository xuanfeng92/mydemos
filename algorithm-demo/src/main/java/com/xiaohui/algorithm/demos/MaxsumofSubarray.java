package com.xiaohui.algorithm.demos;

import com.xiaohui.algorithm.common.Utils;

import java.util.ArrayList;
import java.util.List;

/**
 * 给定一个数组arr，返回子数组的最大累加和
 * 例如，arr = [1, -2, 3, 5, -2, 6, -1]，所有子数组中，[3, 5, -2, 6]可以回累加出最大的和12，所以返12.
 * [要求]
 * 时间复杂度为O(n)O(n)，空间复杂度为O(1)O(1)
 */
public class MaxsumofSubarray {
    public static void main(String[] args) {
        int[] arrays = {1, -2, 3, 5, -2, 6, -1};
        System.out.println("sum:"+ maxsumofSubarray(arrays));
    }

    public static int maxsumofSubarray (int[] arr) {
        if(arr == null || arr.length == 0) return -1;
        int len = arr.length;
        int max = Integer.MIN_VALUE;
        int cur = 0;
        for(int i = 0; i < len; i++){
            cur += arr[i];
            max = Math.max(cur, max);
            // 关键点： cur<0的话，则将前面累加的数组舍去，重新开始累加
            cur = cur < 0 ? 0 : cur;
        }
        return max;
    }
}
