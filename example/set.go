package main

import (
	"github.com/go-redis/redis"
	"github.com/hxun123/ecredis"
	"fmt"
)

// set 操作
func setOperation(client *redis.Client) {
	// 向 blacklist 中添加元素
	client.SAdd("blacklist", "Obama")
	// 再次添加
	client.SAdd("blacklist", "Hillary")
	// 添加新元素
	client.SAdd("blacklist", "the Elder")
	// 向 whitelist 添加元素
	client.SAdd("whitelist", "the Elder")

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist: ", isMember)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("Inter result: ", names)

	// 获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member: ", all)
}

func main() {
	client := ecredis.NewClient("crm.master")
	defer client.Close()

	setOperation(client)

}
