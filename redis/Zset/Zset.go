package Zset

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Zset struct {
	rdb  *redis.Client
	ctx  context.Context
	name string
}

func NewZset(rdb *redis.Client, ctx context.Context, name string) *Zset {
	return &Zset{
		rdb:  rdb,
		ctx:  ctx,
		name: name,
	}
}

func (z *Zset) Add(weight float64, member interface{}) {
	err := z.rdb.ZAdd(z.ctx, z.name, redis.Z{Score: weight, Member: fmt.Sprintf("%v", member)}).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (z *Zset) Remove(member interface{}) {
	err := z.rdb.ZRem(z.ctx, z.name, fmt.Sprintf("%v", member)).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (z *Zset) RangeByScore(min, max string) []redis.Z {
	values, err := z.rdb.ZRangeByScoreWithScores(z.ctx, z.name, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return values
}

func (z *Zset) Card() int64 {
	card, err := z.rdb.ZCard(z.ctx, z.name).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return card
}

func (z *Zset) Score(member interface{}) float64 {
	score, err := z.rdb.ZScore(z.ctx, z.name, fmt.Sprintf("%v", member)).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return score
}

func (z *Zset) IncrBy(increment float64, member interface{}) {
	err := z.rdb.ZIncrBy(z.ctx, z.name, increment, fmt.Sprintf("%v", member)).Err()
	if err != nil {
		fmt.Println(err)
	}
}

/*

#### 4. Sorted Set（有序集合）
有序集合和集合类似，但每个元素都会关联一个分数，Redis 会根据分数对元素进行排序。

- **添加一个或多个元素及其分数到有序集合：**
  ```shell
  ZADD key score member [score member ...]
  ```
- **获取有序集合中的元素数量：**
  ```shell
  ZCARD key
  ```
- **获取有序集合中某个分数范围内的元素：**
  ```shell
  ZRANGE key start stop [WITHSCORES]
  ```
- **移除有序集合中的一个或多个元素：**
  ```shell
  ZREM key member [member ...]
  ```
- **获取某个元素的分数：**
  ```shell
  ZSCORE key member
  ```
- **对有序集合中某个元素的分数进行加减：**
  ```shell
  ZINCRBY key increment member
  ```
*/
