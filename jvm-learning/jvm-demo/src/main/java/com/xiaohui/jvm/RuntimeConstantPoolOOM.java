package com.xiaohui.jvm;

import java.util.ArrayList;
import java.util.List;

public class RuntimeConstantPoolOOM {
    public static void main(String[] args) {
        // 使用List保持常量池引用，笔描Full GC回收常量池行为
        List<String> list = new ArrayList();
        int i =0;
        while (true){
            list.add(String.valueOf(i++).intern());
        }
    }
}
