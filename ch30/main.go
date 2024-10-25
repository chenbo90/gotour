package main

import "fmt"

func main() {

	//ch := make(chan int, 1)
	//for i := 1; i <= 10; i++ {
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x)
	//	case ch <- i:
	//	}
	//}

	m := make(map[string][]string)

	var slice1 []string
	var slice2 []string

	// 使用 append 动态添加元素
	slice1 = append(slice1, "192.168.0.1", "192.168.0.2")
	slice2 = append(slice2, "192.168.0.3", "192.168.0.4")
	m["www.baidu.com"] = slice1
	m["www.google.com"] = slice2

	m2 := make(map[string]string)
	m2["first"] = "hello"
	m2["second"] = "hello"
	fmt.Println(m2)

}
