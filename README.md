# ecredis
基于go-redis进行简单封装，使用hash或者随机返回client，实现了主从分离和增减redis单机实例
## 下载
`go get -u github.com/hxun123/ecredis`
## 使用示例
进入当前包example目录，直接执行如下即可：
```
go run string.go
go run hash.go
go run set.go
go run list.go
go run connectPool.go
```