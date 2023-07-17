package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/users",handleUsers)
	http.HandleFunc("/users1",handleUsers1)
	http.HandleFunc("/users2",handleUsers2)
	http.ListenAndServe("127.0.0.1:8080",nil)
	//time.Sleep(60*time.Second)
}

func handleUsers(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"ID:1,Name:张三")
	fmt.Fprintln(w,"ID:2,Name:李四")
	fmt.Fprintln(w,"ID:3,Name:王五")
}

func handleUsers1(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w,"ID:4,Name:张三")
		fmt.Fprintln(w,"ID:5,Name:李四")
		fmt.Fprintln(w,"ID:6,Name:王五")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w,"not found")
	}

}

func handleUsers2(w http.ResponseWriter,r *http.Request){

	var users = []User{
		{ID:7,Name:"张三"},
		{ID:8,Name:"李四"},
		{ID:9,Name:"王五"},
	}

	switch r.Method{
	case "GET":
		w.WriteHeader(http.StatusOK)
		users,err := json.Marshal(users)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
		}else{
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w,"{\"message:\":\"not found\"}")
	}

}

type User struct {
	ID int
	Name string
}
