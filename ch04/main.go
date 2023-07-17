package main

import "fmt"

func main() {
	fmt.Println("---this is ch04 test---")

	//声明数组
	array := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(array[2])

	//遍历数组
	for k, v := range array {
		fmt.Printf("数组索引：%d 对应值是：%s\n", k, v)
	}

	//遍历数组,舍弃健
	for _, v := range array {
		fmt.Printf("对应值是：%s\n", v)
	}

	//切片
	slice := array[2:5]
	slice[1] = "f"
	fmt.Println(array)
	fmt.Println(slice)

	//切片声明1
	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1), cap(slice1))

	//切片声明2
	slice2 := make([]string, 4, 8)
	fmt.Println(len(slice2), cap(slice2))

	/**
	*总结：Go语言开发中通常会优先选择切片，因为它高效，内存占用小
	 */

}
