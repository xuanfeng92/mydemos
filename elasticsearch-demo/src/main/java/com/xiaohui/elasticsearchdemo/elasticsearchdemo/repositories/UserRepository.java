package com.xiaohui.elasticsearchdemo.elasticsearchdemo.repositories;

import com.xiaohui.elasticsearchdemo.elasticsearchdemo.pojo.User;
import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

public interface UserRepository extends ElasticsearchRepository<User, Long> {
}
