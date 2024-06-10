# docker命令
## 通用
```bash
查看 Docker 版本
docker --version

启动 Docker 服务//windows 不可用
systemctl start docker

停止 Docker 服务 //windows 不可用
systemctl stop docker

重启 Docker 服务//windows 不可用
systemctl restart docker

查看 Docker 服务状态//windows 不可用
systemctl status docker

拉取镜像
docker pull [OPTIONS] NAME[:TAG|@DIGEST]

查看本地镜像
docker images

删除镜像
docker rmi [OPTIONS] IMAGE [IMAGE...]

搜索镜像
docker search [OPTIONS] TERM

运行容器
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]

列出正在运行的容器
docker ps

列出所有容器
docker ps -a

停止容器
docker stop CONTAINER

启动已停止的容器
docker start CONTAINER

重启容器
docker restart CONTAINER

删除容器
docker rm CONTAINER

进入容器
docker exec -it CONTAINER COMMAND

查看容器日志
docker logs CONTAINER

查看容器内部文件系统
docker cp CONTAINER:/src path/to/destination

查看 Docker 信息
docker info

查看 Docker 帮助
docker --help
```
## mysql
```bash
docker run --name test-mysql -v "F:\docker_vh\test-mysql:/var/lib/mysql" -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=Gorm -p 3008:3306 mysql
```


--name some-mysql: 给容器指定一个名称，这里使用了 some-mysql。

-e MYSQL_ROOT_PASSWORD=my-secret-pw: 设置 MySQL root 用户的密码，这里是 my-secret-pw。请确保使用一个强密码。

-d: 表示在后台运行容器。

mysql:tag: 指定要使用的 MySQL 镜像和标签（tag）。如果不指定标签，默认使用 latest

-v /path/to/mysql/data:/var/lib/mysql 表示将宿主机的 /path/to/mysql/data 目录挂载到容器的 /var/lib/mysql 目录，这样 MySQL 数据就会持久化存储在宿主机上。
