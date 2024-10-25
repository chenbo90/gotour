FROM golang:1.16-alpine AS build

WORKDIR /home/app/
COPY ../ /home/app/
#RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go build -o /home/app/work/server main.go

FROM alpine
ENV TZ Asia/Shanghai
COPY --from=build /home/app/work/ /work/

WORKDIR /work
EXPOSE 8080

ENTRYPOINT ["./server"]