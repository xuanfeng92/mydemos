package com.xiaohui.jvm;

import java.util.ArrayList;
import java.util.List;

public class OutOfMemory {
    public static void main(String[] args) {
        List<TestObject> list=new ArrayList<TestObject>();
        while(true){
            list.add(new TestObject());
        }
    }

    static class TestObject{}
}
