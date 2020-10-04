package com.xiaohui.jvm;

/**
 * VM Args: -Xss128k
 */
public class StackOverflowTest {
    private int stackLength = 1;
    public void stackLeak(){
        stackLength ++;
        stackLeak();
    }

    public static void main(String[] args) throws Exception {
        StackOverflowTest stackOverflowTest = new StackOverflowTest();
        try {
            stackOverflowTest.stackLeak();
        } catch (Exception e) {
            System.out.println("stack length:"+ stackOverflowTest.stackLength);
            throw e;
        }
    }
}
