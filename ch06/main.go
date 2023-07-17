package main

import "fmt"

func main() {
	var p person
	p.age = 20
	p.name = "张三"
	fmt.Println(p)

	p1 := person{"李四", 30, address{"江苏省", "南京市"}}
	fmt.Println(p1)

	p2 := person{name: "王五", age: 44, addr: address{"江苏省", "南京市"}}
	fmt.Println(p2.addr.province)
	methodresult := p2.String()
	fmt.Println(methodresult)

	//以值类型接收者实现接口的时候，类型本身和该类型的指针类型都实现了该接口
	printString(p1)

	printString(&p2)

	p3 := NewPerson("赵六")
	fmt.Println(p3)
	printString(p3)

}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

//-----分割线----
//定义接口，存在一个方法String
type Stringer interface {
	String() string
}

//接口的实现
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

// func (p *person) String() string {
// 	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
// }

//定义一个函数，参数是个接口
func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

//工厂函数
func NewPerson(name1 string) *person {
	return &person{name: name1}
}
