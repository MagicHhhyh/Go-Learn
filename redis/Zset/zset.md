
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
