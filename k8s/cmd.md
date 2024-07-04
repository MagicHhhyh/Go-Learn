# kubectl 
1. **获取资源信息**：
    - `kubectl get nodes`：列出集群中的所有节点。
    - `kubectl get pods`：列出当前命名空间下的所有 Pods。
    - `kubectl get services`：列出当前命名空间下的所有服务。
    - `kubectl get deployments`：列出当前命名空间下的所有 Deployments。

2. **查看资源详细信息**：
    - `kubectl describe <resource> <name>`：显示指定资源的详细信息，例如 `kubectl describe pod my-pod`。

3. **创建资源**：
    - `kubectl create deployment <name> --image=<image>`：创建一个新的 Deployment。
    - `kubectl create service clusterip <name> --tcp=<port>:<target-port>`：创建一个新的 ClusterIP 服务。

4. **更新资源**：
    - `kubectl set image deployment/<name> <container>=<new-image>`：更新 Deployment 中的容器镜像。
    - `kubectl scale deployment <name> --replicas=<number>`：调整 Deployment 的副本数。

5. **编辑资源**：
    - `kubectl edit <resource> <name>`：编辑指定资源的定义，例如 `kubectl edit deployment my-deployment`。

6. **删除资源**：
    - `kubectl delete <resource> <name>`：删除指定的资源，例如 `kubectl delete pod my-pod`。

7. **管理配置和密钥**：
    - `kubectl create configmap <name>`：创建一个新的 ConfigMap。
    - `kubectl create secret generic <name>`：创建一个新的 Secret。

8. **查看日志**：
    - `kubectl logs <pod-name>`：获取 Pod 的日志。
    - `kubectl logs -f <pod-name>`：持续跟踪 Pod 的日志输出。

9. **访问 Pod**：
    - `kubectl exec <pod-name> -- <command>`：在指定的 Pod 中执行命令。
    - `kubectl port-forward <pod-name> <local-port>:<remote-port>`：将本地端口转发到 Pod 的端口。

10. **部署应用**：
    - `kubectl apply -f <filename.yaml>`：使用 YAML 文件部署应用。

11. **滚动更新**：
    - `kubectl rollout status deployment/<name>`：检查 Deployment 的滚动更新状态。

12. **查看集群事件**：
    - `kubectl get events`：列出集群中的事件，这有助于诊断问题。

13. **设置资源配额和限制范围**：
    - `kubectl create quota <name> --hard=<resource>=<amount>`：为命名空间设置资源配额。

14. **查看集群状态**：
    - `kubectl cluster-info`：显示集群信息，包括 Kubernetes 服务的版本和状态。

15. **使用不同的上下文**：
    - `kubectl config use-context <context-name>`：切换到不同的 Kubernetes 集群上下文。

