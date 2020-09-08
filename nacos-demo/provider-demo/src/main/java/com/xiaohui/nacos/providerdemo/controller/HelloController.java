package com.xiaohui.nacos.providerdemo.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RefreshScope
public class HelloController {

    @Value("${myConfig}")
    private String nacosConfig;

    @Value("${jdbcUrl}")
    private String jdbcUrl;

    @GetMapping(value = "hello")
    public String hello() {
        return "hello provider config is:【" + nacosConfig + "】" + "jdbcUrl is 【" + jdbcUrl + "】";
    }
}