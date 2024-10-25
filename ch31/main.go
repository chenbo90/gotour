package main

import "fmt"

func reverseInPlace(input *[]int) {
	for i, j := 0, len(*input)-1; i < j; i, j = i+1, j-1 {
		(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
	}
}

// reverseArray 翻转数组的函数
func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func main() {
	// 示例数组
	//array := []int{1, 2, 3, 4, 5}
	//fmt.Println("原始数组:", array)
	//
	//// 翻转数组
	////reverseArray(array)
	//reverseInPlace(&array)
	//fmt.Println("翻转后的数组:", array)

	set := make(map[string]struct{})
	elements := []string{"one", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}

	for _, element := range elements {
		set[element] = struct{}{}
	}

	fmt.Println(set)

}
