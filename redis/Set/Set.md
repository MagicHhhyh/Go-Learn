
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
