package com.xiaohui.algorithm.demos;

import javax.sound.sampled.SourceDataLine;
import java.util.Arrays;

/***
 * 选择排序,插入排序，哈希排序
 */
public class SortTest {

    public static void main(String[] args) {
        Integer[] arrays = {25,23,5,7,8,2,33,21,6,23};

//        inserSort(arrays);
//        System.out.println("插入排序后："+ Arrays.asList(arrays).toString());
//
//        selectSort(arrays);
//        System.out.println("选择排序后："+ Arrays.asList(arrays).toString());

        shellSort(arrays, 3);
        System.out.println("哈希排序后："+ Arrays.asList(arrays).toString());
    }


    // 插入排序
    public static void inserSort(Comparable[] a) {
        // 将a升序排列
        int n = a.length;
        for (int i = 0; i < n; i++) {
            // 将a[j] 插入到a[j-1],a[j-2],a[j-3]中
            for (int j = i; j > 0 && less(a[j], a[j-1]); j--) {
                exch(a, j, j-1);
            }
        }
    }

    // 选择排序
    public static void selectSort(Comparable[] a) {
        int n = a.length;
        for (int i = 0; i < n; i++) {
            // 将a[i] 和a[i+1...N]中最小元素交换
            int min = i; // 最小元素的索引
            for (int j = i+1; j < n; j++) {
                if (less(a[j], a[min])) min = j;
            }
            exch(a, i, min);
        }
    }

    // 哈希排序
    public static void shellSort(Comparable[] a, int k) {
        int n = a.length;

        // 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
        int h = 1;
        while (h < n/k) h = k*h + 1;

        while (h >= 1) {
            // h-sort the array
            for (int i = h; i < n; i++) {
                // 将a[i] 插入到a[i-h],a[i-2*h],a[i-3h]...之中
                for (int j = i; j >= h && less(a[j], a[j-h]); j -= h) {
                    exch(a, j, j-h);
                }
            }
            h /= k;
        }
    }

    private static boolean less(Comparable v, Comparable w) {
        return v.compareTo(w) < 0;
    }

    private static void exch(Object[] a, int i, int j) {
        Object swap = a[i];
        a[i] = a[j];
        a[j] = swap;
    }
}
