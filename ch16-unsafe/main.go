package main

import (
	"fmt"
	"unsafe"
)

func main(){
	//	指针类型转换
	i := 2
	ip := &i

	//var fp *float64 = (*float64)ip //报错，不能转换

	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(*fp)

}
