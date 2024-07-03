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
