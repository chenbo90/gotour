FROM golang:1.16-alpine AS build

WORKDIR /home/app/
COPY . /home/app/
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go build -o /home/app/work/second main.go
COPY ./first /home/app/work/


FROM alpine
ENV TZ Asia/Shanghai
#从刚才构建的二进制文件拷贝到work目录下
COPY --from=build /home/app/work/ /work/
RUN chmod 777 /work/first


WORKDIR /work
EXPOSE 8080


CMD ["/bin/sh", "-c", "./first & ./second"]
#ENTRYPOINT ["./server"]