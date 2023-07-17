package main

import "fmt"

/**
* ch13 	参数传递：值，引用及指针之间的区别

总结：
	在Go语言中，函数的参数传递只有值传递，且传递的实参都是原始数据的一份拷贝
	如拷贝类型是值类型，那在函数中无法修改原始数据
	如拷贝的内容是指针，那就可在函数中修改原始数据
 */
func main() {
	fmt.Println("---test ch13----")

	//当值类型作为接收者实现了某个接口时，它的指针类型也同样实现了该接口
	addr := address{provice:"江苏省",city:"南京市"}
	printString(addr)
	printString(&addr)
	fmt.Printf("main函数：addr的内存地址%p\n",&addr)
	modifyAddress(&addr)
	fmt.Println("the addr is",addr.provice,addr.city)

	p := person{name:"张三",age:18}
	fmt.Printf("main函数：p的内存地址%p\n",&p)
	modifyPerson1(p)
	fmt.Println("1person name:",p.name,"person age:",p.age)
	modifyPerson2(&p)
	fmt.Println("2person name:",p.name,"person age:",p.age)

	//---map
	fmt.Println("-------测试map-------")
	m := make(map[string]int)
	m["树下听雨"] = 20
	fmt.Println("树下听雨的年龄是：",m["树下听雨"])
	modifyMap(m)
	fmt.Println("树下听雨的年龄是：",m["树下听雨"])

	/**
	结果：
	-------测试map-------
	树下听雨的年龄是： 20
	树下听雨的年龄是： 30

	为什么能够改变map的值，因为map的底层是指针
	 */







}

type address struct {
	provice string
	city string
}

func (addr address)String() string{
	return fmt.Sprintf("the addr is %s%s",addr.provice,addr.city)
}


func printString(s fmt.Stringer){
	fmt.Println(s.String())
}

func modifyAddress(addr *address){
	fmt.Printf("modifyAddress函数：addr的内存地址%p\n",&addr)
	addr.provice = "安徽省"
	addr.city="合肥市"
}

func modifyPerson1(p person){
	fmt.Printf("modifyPerson1函数：person的内存地址%p\n",&p)
	p.name = "李四"
	p.age = 20
}

func modifyPerson2(p *person){
	fmt.Printf("modifyPerson2函数：person的内存地址%p\n",&p)
	p.name = "王五"
	p.age = 22
}



type person struct {
	name string
	age int
}

func modifyMap(m map[string]int){
	m["树下听雨"] = 30
}

