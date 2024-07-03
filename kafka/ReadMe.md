### Config
```go
type KqConf struct {
    service.ServiceConf
    Brokers    []string
    Group      string
    Topic      string
    Offset     string `json:",options=first|last,default=last"`
    Conns      int    `json:",default=1"`
    Consumers  int    `json:",default=8"`
    Processors int    `json:",default=8"`
    MinBytes   int    `json:",default=10240"`    // 10K
    MaxBytes   int    `json:",default=10485760"` // 10M
    Username   string `json:",optional"`
    Password   string `json:",optional"`
}
```
Brokers: kafka 的多个 Broker 节点

Group：消费者组

Topic：订阅的 Topic 主题

Offset：如果新的 topic kafka 没有对应的 offset 信息,或者当前的 offset 无效了(历史数据被删除),那么需要指定从头(first)消费还是从尾(last)部消费

Conns: 一个 kafka queue 对应可对应多个 consumer，Conns 对应 kafka queue 数量，可以同时初始化多个 kafka queue，默认只启动一个

Consumers : go-queue 内部是起多个 goroutine 从 kafka 中获取信息写入进程内的 channel，这个参数是控制此处的 goroutine 数量（⚠️ 并不是真正消费时的并发 goroutine 数量）

Processors: 当 Consumers 中的多个 goroutine 将 kafka 消息拉取到进程内部的 channel 后，我们要真正消费消息写入我们自己逻辑，go-queue 内部通过此参数控制当前消费的并发 goroutine 数量

MinBytes: fetch 一次返回的最小字节数,如果不够这个字节数就等待.

MaxBytes: fetch 一次返回的最大字节数,如果第一条消息的大小超过了这个限制仍然会继续拉取保证 consumer 的正常运行.因此并不是一个绝对的配置,消息的大小还需要受到 broker 的message.max.bytes限制,以及 topic 的max.message.bytes的限制

Username: kafka 的账号

Password：kafka 的密码

## 配置重点
Consumers : go-queue 内部是起多个 goroutine 从 kafka 中获取信息写入进程内的 channel，这个参数是控制此处的 goroutine 数量（⚠️ 并不是真正消费时的并发 goroutine 数量）

Processors: 当 Consumers 中的多个 goroutine 将 kafka 消息拉取到进程内部的 channel 后，我们要真正消费消息写入我们自己逻辑，go-queue 内部通过此参数控制当前消费的并发 goroutine 数量