package main

import (
	"github.com/go-redis/redis"
	"github.com/hxun123/ecredis"
	"fmt"
	"sync"
)

// 连接池管理
func connectPool(client *redis.Client) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xyz%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns)
		}()
	}

	wg.Wait()
}

func main() {
	client := ecredis.NewClient("crm.master")
	defer client.Close()

	connectPool(client)

}
