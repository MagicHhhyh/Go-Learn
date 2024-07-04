## docker中启动mongodb
```bash
docker pull mongo
docker run --name test-mongodb -d -p 27017:27017 mongo
```
设置用户名密码，以及初始数据库
```bash
docker run --name test-mongodb -d -p 27017:27017   -e MONGO_INITDB_ROOT_USERNAME=root  -e MONGO_INITDB_ROOT_PASSWORD=root    mongo
```
在 Docker 命令 docker run 中，各个参数的含义如下：

--name test-mongodb：为容器指定一个名称，在这个例子中是 test-mongodb。如果省略该参数，Docker 会为容器生成一个随机名称。

-d：表示以 detached 模式运行容器，即在后台运行容器，不会阻塞当前的命令行或终端。

-p 27017:27017：端口映射参数，格式为 <主机端口>:<容器端口>。在这个例子中，容器的 27017 端口映射到主机的 27017 端口，使得你可以从主机或其他容器通过网络访问 MongoDB。

-e：设置环境变量。在这个例子中，-e MONGO_INITDB_ROOT_USERNAME=root 和 -e MONGO_INITDB_ROOT_PASSWORD=examplePassword 分别设置了 MongoDB 的 root 用户名和密码。

mongo：指定要运行的镜像名称。。

使用您提供的 MongoDB 认证信息和连接参数，您可以按照以下步骤构建 MongoDB 连接字符串，并在 Go 程序中使用 `mongo-go-driver` 连接到 MongoDB 数据库：

登录设置认证
```go
clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%d", host, port)).
		SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
```

```go
package main

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 用户名、密码、主机、端口和数据库名
	username := "root"
	password := "examplePassword"
	host := "172.26.80.1"
	port := 27017
	Dbname := "test"

	// 创建 MongoDB 连接字符串
	// 注意密码中的特殊字符需要进行URL编码，例如 @ 替换为 %40
	// 这里假设密码没有特殊字符，如果有，请替换为相应的URL编码
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", username, password, host, port, Dbname))

	// 创建一个新的客户端并连接到服务器
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// 在这里执行其他数据库操作...

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
```

请注意以下几点：

1. **密码处理**：如果密码中包含特殊字符，如 `@`、`:`、`/` 等，需要进行 URL 编码。例如，如果密码是 `examplePassword`，并且包含 `@`，则应替换为 `%40`。

2. **数据库操作**：在连接成功后，您可以执行数据库操作，如插入、查询、更新和删除文档。

3. **错误处理**：在实际应用中，您应该根据需要添加更详细的错误处理逻辑。

4. **关闭连接**：在完成数据库操作后，确保调用 `client.Disconnect` 来关闭连接。

5. **使用上下文**：`context.TODO()` 用于提供当前的上下文，您可以根据需要使用其他类型的上下文，例如设置超时或取消操作。

6. **安全提示**：在生产环境中，不要在代码中硬编码密码。考虑使用环境变量或安全的配置管理系统来管理敏感信息。

7. **MongoDB 驱动版本**：请确保使用的 MongoDB 驱动版本与 MongoDB 服务器版本兼容。
