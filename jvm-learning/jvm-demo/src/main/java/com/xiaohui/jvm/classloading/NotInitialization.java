package com.xiaohui.jvm.classloading;

/**
 * 被动使用类字段演示一：
 * 通过子类引用父类的静态字段，不会导致子类初始化
 */
class SuperClass {
    static {
        System.out.println("SuperClass init!");
    }
    public static int value = 123;
}
class SubClass extends SuperClass{
    static {
        System.out.println("SubClass init!");
    }
//    static int value = 234;
}

class ConstClass{
    static {
        System.out.println("ConstClass init!");
    }
    public static final String HELLOWORLD = "hello world";
}
/**
 * 非主动类初始化演示
 */
public class NotInitialization{
    public static void main(String[] args){
        // 1.通过子类引用父类的静态字段，不会导致子类初始化
//        System.out.println(SubClass.value);
        // 2. 数组定义来引用类，不会触发此类的初始化
//        SuperClass[] superClasses = new SuperClass[10];
        // 3. 常量在编译阶段会存入调用类的常量池中，本质上并没有直接引用到定义常量的类，因此不会出发定义常量类的初始化
        System.out.println(ConstClass.HELLOWORLD);
    }
}