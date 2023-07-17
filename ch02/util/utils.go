package utils

import "fmt"

//函数大写，表示该函数的权限是public，其他包可以引用,小写表示私有
func Util() {
	fmt.Println("this is util")
}

func init() {
	fmt.Println("--我是utils包内的init1--")
}

func init() {
	fmt.Println("--我是utils包内的init2--")
}

