package main

/**
*函数和方法
 */
import (
	"errors"
	"fmt"
)

type Age uint

func main() {

	result := sum(1, 2)
	fmt.Println(result)

	result1, err := sum1(1, -2)
	if nil == err {
		fmt.Println(result1)
	} else {
		fmt.Println(err)
	}

	result2, _ := sum2(3, 4)
	fmt.Println(result2)

	age := Age(25)
	age.String()

}

func sum(a int, b int) int {
	return a + b
}

//多值返回
func sum1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}

	return a + b, nil
}

//命名返回参数
func sum2(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

func (age Age) String() {
	fmt.Println("the age is ", age)
}
