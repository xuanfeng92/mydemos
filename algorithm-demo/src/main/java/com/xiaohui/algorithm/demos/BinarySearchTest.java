package com.xiaohui.algorithm.demos;

import java.util.Arrays;

/**
 * 请实现有重复数字的有序数组的二分查找。
 * 输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一
 */
public class BinarySearchTest {
    public static void main(String[] args) {
        int[] data = {3,3,4,4,4,5,6,6,6,7,8,8,12,13,15,16,21,21,22,24,24,27,28,32,34,35,35,36,36,39,40,41,41,42,44,
                44,45,45,47,47,47,47,48,48,50,51,51,53,53,53,54,54,54,56,56,57,59,60,60,60,60,61,62,63,65,65,65,65,
                67,67,68,70,71,71,74,75,75,79,81,84,84,86,86,87,90,90,90,90,91,92,93,94,94,94,95,97,97,98,98,99};
        int[] data2 = {1,1,2,3,7,7,7,9,9,10};
        System.out.println("查找到的值的位置："+binarySearch(100, 102, data));
        System.out.println("查找到的值的位置："+binarySearch(10, 2, data2));

    }

    private static int binarySearch(int n, int v, int[] a){
        if (v > a[n - 1]) return n + 1;
        int lo = 0, hi = n;
        while (lo < hi) {
            int mid = lo + (hi - lo) / 2;
            if (a[mid] < v) {
                lo = mid + 1;
            } else hi = mid;
        }
        return lo + 1;
    }

    private static int binarySearch(int[] a, int fromIndex, int toIndex,
                                     int key) {
        int low = fromIndex;
        int high = toIndex - 1;

        while (low < high) {
            int mid =low + (high- low) >>> 1;
            int midVal = a[mid];

            if (midVal < key)
                low = mid + 1;
            else if (midVal > key)
                high = mid;
//            else
//                return mid+1; // key found
        }
        return (low + 1);  // key not found.
    }
}
