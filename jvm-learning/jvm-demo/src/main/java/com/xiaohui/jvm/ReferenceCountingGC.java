package com.xiaohui.jvm;

import java.lang.ref.SoftReference;

/**
 *
 * -XX:+PrintGC 输出GC日志
 * -XX:+PrintGCDetails 输出GC的详细日志
 * -XX:+PrintGCTimeStamps 输出GC的时间戳（以基准时间的形式）
 * -XX:+PrintGCDateStamps 输出GC的时间戳（以日期的形式，如 2013-05-04T21:53:59.234+0800）
 * -XX:+PrintHeapAtGC 在进行GC的前后打印出堆的信息
 * -Xloggc:../logs/gc.log 日志文件的输出路径
 */
public class ReferenceCountingGC {
    public Object instance = null;
    private static final int _1MB = 1024*1024;

    // 这个成员属性唯一的意义就是占点内存，以便能在GC日志中看清楚是否被回收过
    private byte[] bigSize = new byte[2*_1MB];

    public static void testCG() {
        ReferenceCountingGC objectA = new ReferenceCountingGC();
        ReferenceCountingGC objectB = new ReferenceCountingGC();
        objectA.instance = objectB;
        objectB.instance = objectA;

        objectA = null;
        objectB = null;

        // 假设在这行发生GC，objA和objcB能否被回收？
        System.gc();
    }

    public static void main(String[] args) {
        ReferenceCountingGC.testCG();
    }
}
