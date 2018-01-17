# ecredis
基于go-redis进行简单封装，使用hash或者随机返回client，实现了主从分离和增减redis单机实例
## 下载
`go get -u github.com/hxun123/ecredis`
## 代码示例
进入当前包example目录，使用go run命令执行即可看到效果：
* 字符串相关操作：string.go
* 哈希相关操作：hash.go
* 集合相关操作：set.go
* 链表相关操作：list.go
* 连接池测试：connectPool.go
## 使用示例
新建文件redisTest.go，内容如下：
```
package main

import (
	"github.com/hxun123/ecredis"
	"fmt"
)

func main() {
    // 使用自定义配置文件，使用绝对路径，需要按照示例配置格式编写
    configPath := "/Users/huangxun/workspace/go/config/redis.json"

    // 若测试，传空即可
    // configPath := ""

    // 获取 client
    client := ecredis.NewClient("crm.slave", configPath)

    // 字符串 set
    client.Set("name", "test", 0).Err()

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
输出 test 说明成功
## 注意事项
* 示例配置文件连接地址为 localhost:6379，配置文件为空时需开启本地redis服务
* 自定义配置需注意文件地址是否正确，配置的redis服务是否可用