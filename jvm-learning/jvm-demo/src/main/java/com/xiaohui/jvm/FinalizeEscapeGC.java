package com.xiaohui.jvm;

/**
 * vm Args : -XX:+PrintGCDetails
 * 此代码演示两点：
 * 1. 对象可以被GC时自我拯救
 * 2. 这种自救机会只有一次，因为一个对象的finalize()方法最多只能被系统自动调用一次
 */
public class FinalizeEscapeGC {
    public static FinalizeEscapeGC SAVE_HOOK =null;

    public void isAlive(){
        System.out.println("yes, i am still alive :)");
    }

    @Override
    protected void finalize() throws Throwable {
        super.finalize();
        System.out.println("finalize method executed!");
        FinalizeEscapeGC.SAVE_HOOK = this;
    }

    public static void main(String[] args) throws InterruptedException {
        SAVE_HOOK = new FinalizeEscapeGC();

        // 1. 对象第一次成功拯救自己
        SAVE_HOOK = null;
        System.gc();

        // 2. 由于finalize方法优先级别很低，所以暂停0.5秒等待它
        Thread.sleep(500);
        if (SAVE_HOOK != null){
            SAVE_HOOK.isAlive();
        }else{
            System.out.println("no, i am dead :(");
        }

        // 3. 下面这段代码与上面完全相同，但由于已经调用过自身的finalize方法了，再次就不会调用了
        SAVE_HOOK = null;
        System.gc();  // 第二次GC则不会逃脱了

        Thread.sleep(500);
        if (SAVE_HOOK != null){
            SAVE_HOOK.isAlive();
        }else{
            System.out.println("no, i am dead :(");
        }
    }
}
