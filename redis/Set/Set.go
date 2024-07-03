package Set

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Set struct {
	rdb  *redis.Client
	ctx  context.Context
	name string
}

func NewSet(rdb *redis.Client, ctx context.Context, name string) *Set {
	return &Set{
		rdb:  rdb,
		ctx:  ctx,
		name: name,
	}
}

func (s *Set) Add(members ...interface{}) {
	err := s.rdb.SAdd(s.ctx, s.name, members...).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Set) Remove(members ...interface{}) {
	err := s.rdb.SRem(s.ctx, s.name, members...).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Set) IsMember(member interface{}) bool {
	isMember, err := s.rdb.SIsMember(s.ctx, s.name, member).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return isMember
}

func (s *Set) Members() []string {
	members, err := s.rdb.SMembers(s.ctx, s.name).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return members
}

func (s *Set) Card() int64 {
	card, err := s.rdb.SCard(s.ctx, s.name).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return card
}

/*

#### 3. Set（无序集合）
集合是一个无序的字符串集合，不允许重复元素。

- **添加一个或多个元素到集合：**
  ```shell
  SADD key member [member ...]
  ```
- **移除集合中的一个或多个元素：**
  ```shell
  SREM key member [member ...]
  ```
- **判断一个元素是否在集合中：**
  ```shell
  SISMEMBER key member
  ```
- **获取集合中的所有元素：**
  ```shell
  SMEMBERS key
  ```
- **获取集合的元素数量：**
  ```shell
  SCARD key
  ```

*/
