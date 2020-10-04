package com.xiaohui.jvm.classloading;

class DeadLoopClass{
    static {
        /** 如果不加上这个if语句，编译器执行提示“Initializer does not complte normally 拒绝编译”*/
        if(true){
            System.out.println(Thread.currentThread()+"init DeadLoopClass");
            while (true){

            }
        }
    }
}

public class DeadLoopClassTest {
    public static void main(String[] args) {
        Runnable runnable = new Runnable() {
            public void run() {
                System.out.println(Thread.currentThread() + "start");
                DeadLoopClass deadLoopClass = new DeadLoopClass();
                System.out.println(Thread.currentThread() + " run over");
            }
        };
        Thread thread1= new Thread(runnable);
        Thread thread2= new Thread(runnable);
        thread1.start();
        thread2.start();
    }
}
