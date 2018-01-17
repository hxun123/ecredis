package main

import (
	"github.com/go-redis/redis"
	"github.com/hxun123/ecredis"
	"fmt"
	"time"
)

// String 操作
func stringOperation(client *redis.Client) {
	// 第三个参数是过期时间, 如果是0, 则表示没有过期时间
	err := client.Set("name", "Hans", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	// 这里设置过期时间
	err = client.Set("age", "20", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// 自增
	client.Incr("age")
	// 自增
	client.Incr("age")
	// 自减
	client.Decr("age")

	val, err = client.Get("age").Result()
	if err != nil {
		panic(err)
	}
	// age 的值为21
	fmt.Println("age", val)

	// 因为 key "age" 的过期时间是一秒钟, 因此当一秒后, 此 key 会自动被删除了
	time.Sleep(1 * time.Second)
	val, err = client.Get("age").Result()
	if err != nil {
		// 因为 key "age" 已经过期了, 因此会有一个 redis: nil 的错误
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("age", val)
}

func main() {
	client := ecredis.NewClient("crm.master")
	defer client.Close()

	stringOperation(client)

}
