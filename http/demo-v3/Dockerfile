FROM golang:1.16-alpine AS build

WORKDIR /home/app/
COPY . /home/app/
RUN go env -w GOPROXY=https://goproxy.io,direct

#编译第一个进程
RUN go build -o /home/app/work/first ./first/first.go
#编译第二个进程
RUN go build -o /home/app/work/second ./second/second.go


FROM alpine
ENV TZ Asia/Shanghai
#从刚才构建的二进制文件拷贝到work目录下
COPY --from=build /home/app/work/ /work/

WORKDIR /work
#开启多个端口 8081和8082
EXPOSE 8081
EXPOSE 8082


CMD ["/bin/sh", "-c", "./first & ./second"]
