package main

import (
	"fmt"
	"sync"
	"time"
)

/**
*	ch10 Context
 */
var (
	sum   int
	mutex sync.Mutex //互斥锁
)

func main() {

	//开启100个协程让sum+10
	for i := 0; i < 100; i++ {
		//go add(10)
		go add1(10)
	}

	//防止提前退出
	time.Sleep(1 * time.Second)
	fmt.Println("和为：", sum)

}

func add(i int) {
	sum += i
}

func add1(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}
