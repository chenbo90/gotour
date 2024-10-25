# -f参数指定dockerfile文件，后面的路径指向的是上下文的路径，可以和dockerfile不在一个路径下，默认是dockerfile文件所在的路径。
docker build -t httpdemo:v1 -f /root/code/gotour/http/demo1/docker/Dockerfile /root/code/gotour/http/demo1