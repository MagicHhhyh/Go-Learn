package main

import (
	"GoLearn/redis/Zset"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func GetRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.26.80.1:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
func main() {
	rdb := GetRedis()
	Z := Zset.NewZset(rdb, context.Background(), "zset")
	Z.Add(1.0, "ok")
	fmt.Println(Z.Card())
}
