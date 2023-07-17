package main

import (
	"fmt"
	"time"
)

/**
*	ch08 并发基础
 */

func main() {

	goroutine1()

	goroutine2()

}

func goroutine1() {

	go fmt.Println("树下听雨")
	fmt.Println("我是main goroutine")
	time.Sleep(time.Second) //如果不加这段代码，则主线程走完，就不会走协程了。
}

func goroutine2() {
	ch := make(chan string)

	go func() {
		fmt.Println("林中散步")
		ch <- "goroutine完成"
	}()
	fmt.Println("我是main goroutine")
	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)
}
