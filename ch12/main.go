package main

import "fmt"

/**
* ch12 	指针
 */
func main() {
	fmt.Println("---test ch12----")

	//获取指针变量地址
	name := "树下听雨"
	nameP := &name

	fmt.Println("nameP的值是：", nameP)

	//获取指针指向的内存保存的变量值
	nameV := *nameP
	fmt.Println("nameV的值是：", nameV)

	//修改指针指向的变量值
	*nameP = "大帅比，树下听雨"
	fmt.Println("nameP修改后的值是：", *nameP)

	//--------------------
	age := 18
	modifAge1(age)
	fmt.Println("1-age的值为：", age)

	modifAge2(&age)
	fmt.Println("2-age的值为：", age)

}

func modifAge1(age int) {
	age = 20
}

func modifAge2(age *int) {
	*age = 20
}

//对于是否使用指针作为接收者，有以下几点参考
//01	如果接收者类型是map/slice/channel这类引用类型，不使用指针
//02	如果需要修改接收者，那么需要使用指针
//03	如果接收者是比较大的类型，可以考虑使用指针
