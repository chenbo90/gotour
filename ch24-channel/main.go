package main

import "fmt"

/**
无缓冲通道，必须有至少一个接收方才能发送成功，下面这个就会发送失败，造成死锁
fatal error: all goroutines are asleep - deadlock!
*/
//func main() {
//	ch := make(chan int)
//	ch <- 10
//	fmt.Println("发送成功")
//}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//func main() {
//	ch := make(chan int)
//	go recv(ch) // 创建一个 goroutine 从通道接收值
//	ch <- 10
//	fmt.Println("发送成功")
//}

//func main() {
//	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
//	ch <- 10
//	fmt.Println("发送成功")
//}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

//
//func Producer() <-chan int{
//	ch := make(chan int, 2)
//
//	go func() {
//		defer close(ch)
//		for i := 0;i < 10;i++{
//			if i%2 == 1{
//				ch <- i
//			}
//		}
//		//close(ch)
//	}()
//	return ch
//}

func Consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//close(ch)
	//f2(ch)

	//ch := Producer()
	//total := Consumer(ch)
	//fmt.Print("total is:",total)

	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}

}
