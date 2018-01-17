package main

import (
	"github.com/go-redis/redis"
	"github.com/hxun123/ecredis"
	"fmt"
)

// hash 操作
func hashOperation(client *redis.Client) {
	// 向名称为 user_xyz 的 hash 中添加元素 name
	client.HSet("user_xyz", "name", "xyz")
	// 向名称为 user_xyz 的 hash 中添加元素 age
	client.HSet("user_xyz", "age", "18")

	// 批量地向名称为 user_test 的 hash 中添加元素 name 和 age
	client.HMSet("user_test", map[string]interface{}{"name": "test", "age": 20})
	// 批量获取名为 user_test 的 hash 中的指定字段的值.
	fields, err := client.HMGet("user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test: ", fields)

	// 获取名为 user_xyz 的 hash 中的字段个数
	length, err := client.HLen("user_xyz").Result()
	if err != nil {
		panic(err)
	}
	// 字段个数为2
	fmt.Println("field count in user_xyz: ", length)

	// 删除名为 user_test 的 age 字段
	client.HDel("user_test", "age")
	age, err := client.HGet("user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error: %v\n", err)
	} else {
		// 字段个数为2
		fmt.Println("user_test age is: ", age)
	}
}

func main() {
	client := ecredis.NewClient("crm.master")
	defer client.Close()

	hashOperation(client)

}
