package main

import (
	"github.com/go-redis/redis"
	"github.com/hxun123/ecredis"
	"fmt"
)

// list 操作
func listOperation(client *redis.Client) {
	// 在名称为 fruit 的list尾添加一个值为value的元素
	client.RPush("fruit", "apple")
	// 在名称为 fruit 的list头添加一个值为value的 元素
	client.LPush("fruit", "banana")
	// 返回名称为 fruit 的list的长度
	length, err := client.LLen("fruit").Result()
	if err != nil {
		panic(err)
	}
	// 长度为2
	fmt.Println("length: ", length)

	//返回并删除名称为 fruit 的list中的首元素
	value, err := client.LPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)

	// 返回并删除名称为 fruit 的list中的尾元素
	value, err = client.RPop("fruit").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)
}

func main() {
	client := ecredis.NewClient("crm.master")
	defer client.Close()

	listOperation(client)

}
