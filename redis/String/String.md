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