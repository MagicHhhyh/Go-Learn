
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
