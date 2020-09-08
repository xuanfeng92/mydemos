package com.xiaohui.elasticsearchdemo.elasticsearchdemo;

import com.xiaohui.elasticsearchdemo.elasticsearchdemo.pojo.User;
import com.xiaohui.elasticsearchdemo.elasticsearchdemo.repositories.UserRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.elasticsearch.core.ElasticsearchRestTemplate;
import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;


@SpringBootTest
class ElasticsearchDemoApplicationTests {

    @Autowired
    ElasticsearchRestTemplate elasticsearchRestTemplate;

    @Autowired
    UserRepository userRepository;

    @Test
    void contextLoads() {
        this.elasticsearchRestTemplate.createIndex(User.class);
        this.elasticsearchRestTemplate.putMapping(User.class);
    }

    @Test
    void testAdd(){
//        this.userRepository.save(new User(1l, "中国制造，世界一流",102, "123456"));
        List<User>  users = Arrays.asList(
                new User(2l, "范冰冰",30, "123456"),
                new User(3l, "韩红",45, "zxcvb123"),
                new User(4l, "小辉",25, "zxcvb123"),
                new User(5l, "呵呵呵呵哒",23, "123456")

        );
        this.userRepository.saveAll(users);
    }

    @Test
    void testFind(){
        System.out.println(this.userRepository.findById(3l).get());
    }
}
