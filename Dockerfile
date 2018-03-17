# 官方Golang Image
FROM golang:alpine
LABEL maintainer="boyuan liang(lby89757@hotmail.com)"

ARG home=/go/src/github.com/branLiang/gin-playground
WORKDIR $home
COPY . $home

# 编译并且生成可执行文件
RUN go install -v

# 运行程序
CMD ["gin-playground"]
