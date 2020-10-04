package com.xiaohui.jvm;

/**
 * vm 参数： -verbose:gc -Xms20M -Xmx20M -Xmn10M -XX:+PrintGCDetails -XX:SurvivorRatio=8 -XX:PretenureSizeThreshold=3145728  -XX:+UseConcMarkSweepGC
 */
public class PretenureSizeThresholdTest {
    private static final int _1MB = 1024*1024;
    public static void main(String[] args) {
        byte[] allocation;
        allocation = new byte[6*_1MB];
    }
}
