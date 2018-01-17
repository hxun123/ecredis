# ecredis
基于go-redis进行简单封装，使用hash或者随机返回client，实现了主从分离和增减redis单机实例
## 下载
`go get -u github.com/hxun123/ecredis`
## 代码示例
进入当前包example目录，使用go run命令执行即可看到效果：
```
* string.go 字符串相关操作
* hash.go 哈希相关操作
* set.go 集合相关操作
* list.go 链表相关操作
* connectPool.go 连接池测试
```
## 使用示例
新建文件redisTest.go，内容如下：
```
package main

import "github.com/hxun123/ecredis"

func main() {
    // 使用自定义配置文件，使用绝对路径，需要按照示例配置编写
    configPath := "/Users/huangxun/workspace/go/config/redis.json"

    // 若测试，传空即可
    // configPath := ""

    // 获取client
    client := ecredis.NewClient("crm.slave", configPath)

    // 字符串 get
    name, err := client.Get("name").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(name)
}
```
执行如下：
```
go run redisTest.go
```