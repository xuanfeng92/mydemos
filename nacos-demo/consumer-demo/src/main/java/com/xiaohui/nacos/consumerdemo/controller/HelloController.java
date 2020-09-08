package com.xiaohui.nacos.consumerdemo.controller;

import com.xiaohui.nacos.consumerdemo.feign.ProviderClient;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;

@RestController
public class HelloController {

    @Autowired
    ProviderClient providerClient;

    @GetMapping("hi")
    public String hello(HttpServletRequest request){
        String hello = "hello consumer to "+ providerClient.hello();
        int host = request.getServerPort();  //这里用于区分负载均衡的不同端口
        return hello + "running port is 【"+host+"】";
    }
}
