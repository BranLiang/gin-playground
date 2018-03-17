# 官方Golang Image
FROM golang:alpine
LABEL maintainer="boyuan liang(lby89757@hotmail.com)"

# 定义工作环境位置, 因为代码仓库在github上面所以配置如下
ARG home=/go/src/github.com/branLiang/gin-playground
WORKDIR $home
COPY . $home

# 安装环境依赖
RUN go get -v github.com/gin-gonic/gin

# 编译并且生成可执行文件
RUN go install -v

# 运行程序
CMD ["gin-playground"]
