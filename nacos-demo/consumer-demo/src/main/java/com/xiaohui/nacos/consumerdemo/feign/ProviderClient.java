package com.xiaohui.nacos.consumerdemo.feign;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;

@FeignClient(value = "service-provider")
public interface ProviderClient {

    @GetMapping("/hello")
    public String hello();

}
