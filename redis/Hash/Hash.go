package Hash

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Hash struct {
	rdb  *redis.Client
	ctx  context.Context
	name string
}

func NewHash(rdb *redis.Client, ctx context.Context, name string) *Hash {
	return &Hash{
		rdb:  rdb,
		ctx:  ctx,
		name: name,
	}
}

func (h *Hash) Set(field string, value interface{}) {
	err := h.rdb.HSet(h.ctx, h.name, field, value).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Hash) Get(field string) string {
	val, err := h.rdb.HGet(h.ctx, h.name, field).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func (h *Hash) Delete(fields ...string) {
	err := h.rdb.HDel(h.ctx, h.name, fields...).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Hash) GetAll() map[string]string {
	vals, err := h.rdb.HGetAll(h.ctx, h.name).Result()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return vals
}

func (h *Hash) Length() int64 {
	length, err := h.rdb.HLen(h.ctx, h.name).Result()
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return length
}

func (h *Hash) Exists(field string) bool {
	exists, err := h.rdb.HExists(h.ctx, h.name, field).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return exists
}
