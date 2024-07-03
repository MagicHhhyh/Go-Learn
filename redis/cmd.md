### Redis 数据结构及其 redis-cli 命令

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
- **减小指定整数值：**
  ```shell
  DECRBY key decrement
  ```

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

#### 5. Hash（哈希）
哈希是一个键值对集合，适合存储对象。

- **设置哈希表中的字段：**
  ```shell
  HSET key field value
  ```
- **获取哈希表中的字段值：**
  ```shell
  HGET key field
  ```
- **删除哈希表中的一个或多个字段：**
  ```shell
  HDEL key field [field ...]
  ```
- **获取哈希表中的所有字段和值：**
  ```shell
  HGETALL key
  ```
- **获取哈希表中的字段数量：**
  ```shell
  HLEN key
  ```
- **判断哈希表中是否存在某个字段：**
  ```shell
  HEXISTS key field
  ```
