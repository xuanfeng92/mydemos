package com.xiaohui.jvm.classloading;

import java.io.IOException;
import java.io.InputStream;

public class ClassLoaderTest {
    public static void main(String[] args) throws ClassNotFoundException, IllegalAccessException, InstantiationException {
        ClassLoader myLoader = new ClassLoader() {
            @Override
            public Class<?> loadClass(String name) throws ClassNotFoundException {
                try {
                    String fileName = name.substring(name.lastIndexOf(".") + 1) + ".class";
                    InputStream is = getClass().getResourceAsStream(fileName);
                    if(is == null){
                        return super.loadClass(name);
                    }
                    byte[] bytes = new byte[is.available()];
                    is.read(bytes);
                    return defineClass(name, bytes, 0, bytes.length);
                } catch (IOException e) {
                    e.printStackTrace();
                    throw new ClassNotFoundException(name);
                }
            }

            @Override
            protected Class<?> findClass(String name) throws ClassNotFoundException {
                return super.findClass(name);
            }
        };

        Class<?> obj = myLoader.loadClass("com.xiaohui.jvm.classloading.ClassLoaderTest");
        System.out.println(obj.getClassLoader());
        System.out.println(ClassLoaderTest.class.getClassLoader());

        Object o = obj.newInstance();
        Object o2 = new ClassLoaderTest();

        System.out.println(o instanceof com.xiaohui.jvm.classloading.ClassLoaderTest);
        System.out.println(o2 instanceof com.xiaohui.jvm.classloading.ClassLoaderTest);
    }
}
