package ecredis

import (
	"crypto/rand"
	"github.com/go-redis/redis"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"hash/crc32"
	"math/big"
	"strings"
)

// 定义配置结构
type conf struct {
	Addr     string
	Password string
	Db       int
}

// 获取配置
func getConfig(key string) conf {
	// 定义配置变量
	var config map[string]map[string][]conf

	// 从文件读取json, todo $GOPATH 替换为真实目录
	jsonData, err := ioutil.ReadFile("$GOPATH/github.com/hxun123/ecredis/config/redis.json")
	if err != nil {
		fmt.Println(err)
	}

	// json解析
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		fmt.Println(err)
	}

	// 获取哈希配置项
	var keyOne, keyTwo string
	keyArr := strings.Split(key, ".")
	if len(keyArr) == 2 {
		keyOne = keyArr[0]
		keyTwo = keyArr[1]
	} else {
		keyOne = "default"
		keyTwo = "master"
	}
	configArr := config[keyOne][keyTwo]

	// 获取配置下标
	i := getRandMod(key, len(configArr))

	// 返回配置项
	return configArr[i]
}

// 获取哈希配置项
func getCrc32Mod(key string, length int) int {
	// 根据key获取crc32值
	data := []byte(key)
	res := crc32.ChecksumIEEE(data)

	// 根据配置项个数取模
	return int(res) % length
}

// 获取随机配置项
func getRandMod(key string, length int) int {
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		fmt.Println(err)
	}

	return int(rnd.Int64())
}

// 新建client
func NewClient(key string) *redis.Client {
	// 根据key获取配置信息
	config := getConfig(key)
	addr := config.Addr
	password := config.Password
	db := config.Db

	// 设置连接池的大小为1000
	poolSize := 1000

	// 创建client
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: poolSize,
	})

	// 测试是否有效
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
