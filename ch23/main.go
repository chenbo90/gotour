package main

import (
	"fmt"
	"unicode"
)

//修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	fmt.Printf("字符串长度为：%v",len(s))
	fmt.Println()
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

//统计字符串中中文汉字的个数
func  countHan(s string) int{
	fmt.Println("输入的汉字是：",s)
	var count int
	for _,c :=range s{
		if unicode.Is(unicode.Han,c){
			count ++
		}
	}
	return count
}

//找到一组数字中，只出现一次的数字
func findSingleNum (nums []int) int{
	res := 0
	for _,num := range nums{
		res ^= num
	}
	fmt.Println("结果为：",res)
	return res
}

func main(){
	changeString()

	traversalString()

	s := "Hello我是小王子"
	han := countHan(s)
	fmt.Println("汉字的个数为：",han)

	num := []int{1, 3, 5, 7, 9, 11, 3, 1, 9, 5, 11}
	findSingleNum(num)
}
