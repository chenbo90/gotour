package main

import "fmt"

func run(){
	fmt.Println("hi")

	defer func() {
		fmt.Println("最后执行----")
	}()

	fmt.Println("hello")
}

func main(){
	run()
}

