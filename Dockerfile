# 选择构建用基础镜像。更新为更高版本的Go以兼容依赖库。根据您的需求，您可以在 Docker Hub 上找到更合适的版本。
FROM golang:1.20-alpine as builder

# 指定构建过程中的工作目录
WORKDIR /app

# 将当前目录（dockerfile所在目录）下所有文件都拷贝到工作目录下（.dockerignore中文件除外）
COPY . /app/

# 执行代码编译命令。操作系统参数为linux，编译后的二进制产物命名为main，并存放在当前目录下。
RUN GOOS=linux go build -o main .

# 选用运行时所用基础镜像。为了减小镜像大小，选用了alpine的较新版本。
FROM alpine:3.15

# 容器默认时区为UTC，如需使用上海时间请启用以下时区设置命令
# RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo Asia/Shanghai > /etc/timezone

# 使用 HTTPS 协议访问容器云调用证书安装
RUN apk add --no-cache ca-certificates

# 指定运行时的工作目录
WORKDIR /app

# 将构建产物/app/main拷贝到运行时的工作目录中
COPY --from=builder /app/main /app/

# 执行启动命令
CMD ["/app/main"]