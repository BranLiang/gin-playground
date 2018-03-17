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
RUN go install $home

# 当container运行时，默认执行编译后文件
ENTRYPOINT /go/bin/gin-playground

# 对外开放接口8080，纯文字提示这里并不真正开放接口！
EXPOSE 8080
