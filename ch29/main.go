package main

import (
	"fmt"
	//"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

var ch1 = make(chan int, 1)

func main() {

	// 启动两个goroutine来向channel发送数据

	go request()
	go response()

	fmt.Println("i am main.")
	time.Sleep(10 * time.Second)
}

func request() {
	//time.Sleep(1 * time.Second)
	go func() {
		select {
		// 当数据库调用完毕则执行取出
		case x := <-ch1:
			// 假如此300毫秒先到了，而readDB()还没有执行完毕则返回超时信息。
			// 300ms => 此实践中并不会触发超时，这是由于我们模拟的数据库读取还是比较简单的。
			// 此处使用 100s 来验证超时
			fmt.Println("recieve channel value is :%d ", x)
		case <-time.After(1000 * time.Millisecond):
			resp := "{\"err\":\"请求超时.\"}"
			fmt.Println(resp)
		}
	}()
	fmt.Println("i am request")
}

func response() {
	time.Sleep(5 * time.Second)
	ch1 <- 200
	fmt.Println("i am response")
}
