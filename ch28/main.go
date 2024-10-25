package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	//var s = "{\"personList\":[{\"age\":30,\"name\":\"chenbo\"},{\"age\":3,\"name\":\"pipi\"}]}"
	//var pp PersonList
	//err1 := json.Unmarshal([]byte(s), &pp) // 将[]byte转换为JSON对象
	//
	//if err1 != nil {
	//	fmt.Println(err1)
	//} else {
	//	fmt.Println(pp.personList[0])
	//}

	var s = "{\"personList\":[{\"age\":30,\"name\":\"chenbo\"},{\"age\":3,\"name\":\"pipi\"}]}"
	var pl PersonList
	err := json.Unmarshal([]byte(s), &pl)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("%+v\n", pl)

	//jsonStr := `[{"name":"Alice","age":25}, {"name":"Bob","age":30}]`
	//
	//var persons []Person // 创建空的对象切片
	//
	//err := json.Unmarshal([]byte(jsonStr), &persons)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, person := range persons {
	//	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	//}

	var haha map[string]string
	xixi := make(map[string]string)
	fmt.Println(haha)
	fmt.Println(xixi)
	s1 := "/get?name=bobo"
	split := strings.Split(s1, "?")
	fmt.Println(split)

	//a := 10
	//b := 5
	//c := 2
	//f := float64(b) / a
	//fmt.Println(f)
	//fmt.Println(c / a)

	// 定义两个整数
	//dividend := 10
	//divisor := 3

	// 进行除法运算，由于divisor是float64，所以结果也会是float64
	//quotient := float64(float64(dividend) / divisor)

	// 打印结果，默认会保留小数点后很多位
	//fmt.Println(quotient) // 输出可能是 3.3333333333333335

	aa := 25
	bb := 84
	proportion := float64(aa) / float64(bb)
	fmt.Println(proportion)
	//floatVal, _ := decimal.NewFromFloat(proportion * 100).RoundFloor(2).Float64()
	n := int64(40)
	breakerMessage := fmt.Sprintf("触发熔断，熔断时间是 %v %s", n, " s")
	fmt.Println(breakerMessage)

	current := time.Now().UnixNano()
	var lastRefillNanoSec int64
	if current > lastRefillNanoSec+2e9 {
		lastRefillNanoSec = current
	}

	sec := float64(lastRefillNanoSec+2.5e9-current) / 1000000000
	fmt.Println(sec)

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//type PersonList []Person

type PersonList struct {
	PersonList []Person `json:"personList"`
}
