package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter, r *http.Request){
	//解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v",err)
	}
	u1 := User{
		Name:"小王子",
		Gender:"男",
		Age: 20,
	}

	m1 := map[string] interface{}{
		"Name": "大王子",
		"Gender": "男",
		"Age": 30,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	//渲染模板
	//t.Execute(w,u1)
	//t.Execute(w,m1)
	t.Execute(w,map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})

}

func main(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err:%v",err)
	}
}
