# 使用golang镜像作为基础镜像
FROM golang:latest

# 安装所需依赖
RUN apt-get update && \
    apt-get install -y git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# 设置工作目录
WORKDIR /site

# 将本地代码复制到容器中
COPY . .

# 下载相关依赖
RUN go mod download

# 启动应用程序
CMD ["go", "run", "./bin/main.go"]