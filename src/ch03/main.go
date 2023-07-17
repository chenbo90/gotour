package main

import "fmt"

func main() {
	fmt.Println("---this is ch03---")

	testIf()

	testSwitch()

	testFor()
}

func testIf() {
	i := 10

	if i > 10 {
		fmt.Println("i > 10")
	} else {
		fmt.Println("i <= 10")
	}

}

func testSwitch() {
	switch j := 1; j {
	case 1:
		fmt.Println("1匹配上了")
		fallthrough //继续向下执行
	case 2:
		fmt.Println("1匹配上了，也执行了分支2")
		fallthrough
	default:
		fmt.Println("没有匹配")
	}
}

func testFor() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println("累加之和结果是：", sum)

}
