
### Deployment 

**Deployment** 是用来声明式地管理无状态应用的声明周期的资源。它确保指定数量的 Pod 副本始终运行，并且可以在 Pod 模板更新时进行滚动更新。

- **`apiVersion`**: 指定使用的 Kubernetes API 版本，对于 Deployment 通常是 `apps/v1`。
- **`kind`**: 资源类型，Deployment 表示这是一个 Deployment 资源。
- **`metadata`**: 包含资源的元数据，如名称和标签。
    - **`name`**: 资源的名称。
    - **`labels`**: 用于逻辑分组和选择器匹配的标签。
- **`spec`**: 定义 Deployment 的期望状态。
    - **`replicas`**: 需要运行的 Pod 副本数量。
    - **`selector`**: 定义如何匹配 Pod 模板中的标签。
    - **`template`**: Pod 模板，定义了 Pod 的规格。
        - **`metadata`**: Pod 的元数据，如标签。
        - **`spec`**: Pod 的规格，包括容器列表。
            - **`containers`**: 容器列表，每个容器包含以下信息：
                - **`name`**: 容器的名称。
                - **`image`**: 容器使用的 Docker 镜像。
                - **`ports`**: 容器要暴露的端口列表，每个端口包含：
                    - **`containerPort`**: 容器内部监听的端口号。

### Service

**Service** 定义了访问一组 Pod 的策略，无论这些 Pod 如何变化。Service 为 Pod 提供一个统一的接口和负载均衡。

- **`apiVersion`**: 指定使用的 Kubernetes API 版本，对于 Service 通常是 `v1`。
- **`kind`**: 资源类型，Service 表示这是一个 Service 资源。
- **`metadata`**: 包含资源的元数据，如名称。
    - **`name`**: 资源的名称。
- **`spec`**: 定义 Service 的期望状态。
    - **`selector`**: 定义 Service 如何选择要代理的 Pods，通常是基于标签选择。
    - **`ports`**: 定义 Service 监听的端口列表，每个端口包含：
        - **`protocol`**: 端口使用的协议，通常是 TCP 或 UDP。
        - **`port`**: 服务监听的端口，这是外部访问 Service 时使用的端口。
        - **`targetPort`**: 要转发到 Pod 的端口，可以是数字或名称。
    - **`type`**: Service 的类型，常见的有：
        - **`ClusterIP`**: 默认类型，Service 仅在集群内部可访问。
        - **`NodePort`**: 在集群的所有节点上暴露一个端口，可以通过节点 IP 和该端口访问 Service。
        - **`LoadBalancer`**: 在云服务提供商上，分配一个外部可访问的负载均衡器。

### 选择器（Selector）

选择器是一组用于匹配 Pods 的标签。Service 使用选择器来确定哪些 Pods 应该接收通过 Service 发送的流量。选择器必须与 Pod 模板中的标签匹配，以便 Service 能够正确地将流量转发到这些 Pods。

### 端口（Ports）

在 Service 中，端口定义了 Service 监听的网络端口，以及如何将这些端口映射到 Pods。`port` 是外部访问 Service 时使用的端口，而 `targetPort` 是 Service 将流量转发到 Pod 的实际端口，它可以是一个端口号或在 Pod 定义中定义的端口名称。

