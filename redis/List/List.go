package List

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type List struct {
	rdb  *redis.Client
	ctx  context.Context
	name string
}

func NewList(rdb *redis.Client, ctx context.Context, name string) *List {
	return &List{
		rdb:  rdb,
		ctx:  ctx,
		name: name,
	}
}

func (l *List) LPush(values ...interface{}) {
	err := l.rdb.LPush(l.ctx, l.name, values...).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (l *List) RPush(values ...interface{}) {
	err := l.rdb.RPush(l.ctx, l.name, values...).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (l *List) LPop() string {
	val, err := l.rdb.LPop(l.ctx, l.name).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (l *List) RPop() string {
	val, err := l.rdb.RPop(l.ctx, l.name).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (l *List) LIndex(index int64) string {
	val, err := l.rdb.LIndex(l.ctx, l.name, index).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (l *List) LLen() int64 {
	length, err := l.rdb.LLen(l.ctx, l.name).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return length
}

func (l *List) LRange(start, stop int64) []string {
	values, err := l.rdb.LRange(l.ctx, l.name, start, stop).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return values
}

/*
#### 2. List（列表）
列表是按顺序排列的字符串集合，可以从左或右推入和弹出元素。

- **在列表左边推入一个或多个值：**
  ```shell
  LPUSH key value [value ...]
  ```
- **在列表右边推入一个或多个值：**
  ```shell
  RPUSH key value [value ...]
  ```
- **从列表左边弹出一个值：**
  ```shell
  LPOP key
  ```
- **从列表右边弹出一个值：**
  ```shell
  RPOP key
  ```
- **获取列表中的一个元素：**
  ```shell
  LINDEX key index
  ```
- **获取列表的长度：**
  ```shell
  LLEN key
  ```
- **获取列表中的一段元素：**
  ```shell
  LRANGE key start stop
  ```
*/
