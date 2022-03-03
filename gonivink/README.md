利用 Docker 环境利用 go-zero 快速开发一个商城系统，让你快速上手微服务。
## 1 环境要求

Golang 1.16+
Etcd
Redis
Mysql
Prometheus
Grafana
Jaeger
DTM

## 2 Docker 本地开发环境搭建

为了方便开发调试，我们使用 Docker 构建本地开发环境。Windows 和 macOS 系统可下载 Docker Desktop 安装使用，具体下载安装方法可自行搜索相关教程。
这里我们使用 Docker Compose 来编排管理我们的容器，创建如下目录：
gonivinck
├── dtm                   # DTM 分布式事务管理器
│   ├── config.yml        # DTM 配置文件
│   └── Dockerfile
├── etcd                  # Etcd 服务注册发现
│   └── Dockerfile
├── golang                # Golang 运行环境
│   └── Dockerfile
├── grafana               # Grafana 可视化数据监控
│   └── Dockerfile
├── jaeger                # Jaeger 链路追踪
│   └── Dockerfile
├── mysql                 # Mysql 服务
│   └── Dockerfile
├── mysql-manage          # Mysql 可视化管理
│   └── Dockerfile
├── prometheus            # Prometheus 服务监控
│   ├── Dockerfile
│   └── prometheus.yml    # Prometheus 配置文件
├── redis                 # Redis 服务
│   └── Dockerfile
├── redis-manage          # Redis 可视化管理
│   └── Dockerfile
├── .env                  # env 配置
└── docker-compose.yml

### 2.1 编写 Dockerfile
golang 容器的 Dockerfile 代码:
```dockerfile
FROM golang:1.17

LABEL maintainer="Ving <ving@nivin.cn>"

ENV GOPROXY https://goproxy.cn,direct

# 安装必要的软件包和依赖包
USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim 

# 安装 goctl
USER root
RUN GOPROXY=https://goproxy.cn/,direct go install github.com/tal-tech/go-zero/tools/goctl@cli

# 安装 protoc
USER root
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc $GOPATH/bin

# 安装 protoc-gen-go
USER root
RUN go get -u github.com/golang/protobuf/protoc-gen-go@v1.4.0

# $GOPATH/bin添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR /usr/src/code

EXPOSE 8000
EXPOSE 8001
EXPOSE 8002
EXPOSE 8003
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 9003
```
### 2.2 编写 .env 配置文件
```env
# 设置时区
TZ=Asia/Shanghai
# 设置网络模式
NETWORKS_DRIVER=bridge


# PATHS ##########################################
# 宿主机上代码存放的目录路径
CODE_PATH_HOST=./code
# 宿主机上Mysql Reids数据存放的目录路径
DATA_PATH_HOST=./data


# MYSQL ##########################################
# Mysql 服务映射宿主机端口号，可在宿主机127.0.0.1:3306访问
MYSQL_PORT=3306
MYSQL_USERNAME=admin
MYSQL_PASSWORD=123456
MYSQL_ROOT_PASSWORD=123456

# Mysql 可视化管理用户名称，同 MYSQL_USERNAME
MYSQL_MANAGE_USERNAME=admin
# Mysql 可视化管理用户密码，同 MYSQL_PASSWORD
MYSQL_MANAGE_PASSWORD=123456
# Mysql 可视化管理ROOT用户密码，同 MYSQL_ROOT_PASSWORD
MYSQL_MANAGE_ROOT_PASSWORD=123456
# Mysql 服务地址
MYSQL_MANAGE_CONNECT_HOST=mysql
# Mysql 服务端口号
MYSQL_MANAGE_CONNECT_PORT=3306
# Mysql 可视化管理映射宿主机端口号，可在宿主机127.0.0.1:1000访问
MYSQL_MANAGE_PORT=1000


# REDIS ##########################################
# Redis 服务映射宿主机端口号，可在宿主机127.0.0.1:6379访问
REDIS_PORT=6379

# Redis 可视化管理用户名称
REDIS_MANAGE_USERNAME=admin
# Redis 可视化管理用户密码
REDIS_MANAGE_PASSWORD=123456
# Redis 服务地址
REDIS_MANAGE_CONNECT_HOST=redis
# Redis 服务端口号
REDIS_MANAGE_CONNECT_PORT=6379
# Redis 可视化管理映射宿主机端口号，可在宿主机127.0.0.1:2000访问
REDIS_MANAGE_PORT=2000


# ETCD ###########################################
# Etcd 服务映射宿主机端口号，可在宿主机127.0.0.1:2379访问
ETCD_PORT=2379


# PROMETHEUS #####################################
# Prometheus 服务映射宿主机端口号，可在宿主机127.0.0.1:3000访问
PROMETHEUS_PORT=3000


# GRAFANA ########################################
# Grafana 服务映射宿主机端口号，可在宿主机127.0.0.1:4000访问
GRAFANA_PORT=4000


# JAEGER #########################################
# Jaeger 服务映射宿主机端口号，可在宿主机127.0.0.1:5000访问
JAEGER_PORT=5000


# DTM #########################################
# DTM HTTP 协议端口号
DTM_HTTP_PORT=36789
# DTM gRPC 协议端口号
DTM_GRPC_PORT=36790
```
### 2.3 编写 docker-compose.yml 配置文件
```dockerfile
version: '3.5'
# 网络配置
networks:
  backend:
    driver: ${NETWORKS_DRIVER}

# 服务容器配置
services:
  golang:                                # 自定义容器名称
    build:
      context: ./golang                  # 指定构建使用的 Dockerfile 文件
    environment:                         # 设置环境变量
      - TZ=${TZ}
    volumes:                             # 设置挂载目录
      - ${CODE_PATH_HOST}:/usr/src/code  # 引用 .env 配置中 CODE_PATH_HOST 变量，将宿主机上代码存放的目录挂载到容器中 /usr/src/code 目录
    ports:                               # 设置端口映射
      - "8000:8000"
      - "8001:8001"
      - "8002:8002"
      - "8003:8003"
      - "9000:9000"
      - "9001:9001"
      - "9002:9002"
      - "9003:9003"
    stdin_open: true                     # 打开标准输入，可以接受外部输入
    tty: true
    networks:
      - backend
    restart: always                      # 指定容器退出后的重启策略为始终重启

```
### 2.4 构建与运行
```
docker-compose up -d
```
