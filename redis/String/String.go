package String

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type String struct {
	rdb  *redis.Client
	ctx  context.Context
	name string
}

func NewString(rdb *redis.Client, ctx context.Context, name string) *String {
	return &String{
		rdb:  rdb,
		ctx:  ctx,
		name: name,
	}
}

func (s *String) Set(value string) {
	err := s.rdb.Set(s.ctx, s.name, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *String) Get() string {
	val, err := s.rdb.Get(s.ctx, s.name).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (s *String) Increment() {
	err := s.rdb.Incr(s.ctx, s.name).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *String) IncrementBy(increment int64) {
	err := s.rdb.IncrBy(s.ctx, s.name, increment).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *String) Decrement() {
	err := s.rdb.Decr(s.ctx, s.name).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *String) DecrementBy(decrement int64) {
	err := s.rdb.DecrBy(s.ctx, s.name, decrement).Err()
	if err != nil {
		fmt.Println(err)
	}
}

/*
#### 1. String（字符串）
字符串是 Redis 最基本的数据类型，可以存储任何形式的字符串。

- **设置键值对：**
  ```shell
  SET key value
  ```
- **获取键的值：**
  ```shell
  GET key
  ```
- **设置键的值并返回旧值：**
  ```shell
  GETSET key value
  ```
- **增加整数值：**
  ```shell
  INCR key
  ```
- **增加指定整数值：**
  ```shell
  INCRBY key increment
  ```
- **减小整数值：**
  ```shell
  DECR key
  ```
*/
