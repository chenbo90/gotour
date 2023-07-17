package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)
type person struct {
	Name string `json:"name"` //json串就可以使用小写
	Age int `json:"age"`
}

func (p person)String() string{
	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
}

func main(){
	i := 3
	iv:=reflect.ValueOf(i)
	it:=reflect.TypeOf(i)
	fmt.Println(iv,it)

	//互转
	iv0 := iv.Interface().(int)
	fmt.Println(iv0)

	//修改值，传递一个指针
	iv1 := reflect.ValueOf(&i)
	iv1.Elem().SetInt(4)
	fmt.Println(i)

	//
	p := person{Name : "树下听雨",Age:32}
	ppv :=reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("陈波")

	fmt.Println("修改后的值：",p)

	ppt := reflect.TypeOf(p)
	
	//遍历字段
	for i :=0; i < ppt.NumField();i++  {
		fmt.Println("字段：",ppt.Field(i).Name)
	}
	//遍历字段中key为json的tag
	for i :=0; i < ppt.NumField();i++  {
		sf := ppt.Field(i)
		fmt.Printf("字段%s上，json tag为 %s\n",sf.Name,sf.Tag.Get("json"))
	}

	//遍历方法
	for i :=0; i < ppt.NumMethod();i++  {
		fmt.Println("方法：",ppt.Method(i).Name)
	}

	//断言的方式判断是否实现了某个接口
	isImplements(p)

	//反射的方式判断是否实现了某个接口
	stringType:=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer",ppt.Implements(stringType))

	//----------
	fmt.Println("--------JSON和结构体之间的互转--------")
	p1 := person{Name:"张三丰",Age:99}
	jsonB,err:=json.Marshal(p1)
	if err == nil{
		fmt.Println(string(jsonB))
	}

	//json to struct
	respJSON := "{\"Name\":\"白眉鹰王\",\"Age\":88}"
	json.Unmarshal([]byte(respJSON),&p1)
	fmt.Println(p1)

	respJSON1 := "{\"name\":\"金毛狮王\",\"age\":66}" //结构体字段使用tag，那么json字符串就可以使用小写，最后转换成结构体的大写
	json.Unmarshal([]byte(respJSON1),&p1)
	fmt.Println(p1)

}

func isImplements(v interface{}){
	if _,ok :=v.(fmt.Stringer);ok{
		fmt.Printf("实现了%s接口","Stringer")
	}
}



