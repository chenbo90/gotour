package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("ginusers",listUser)
	r.Run("localhost:8080")
}

func listUser(c *gin.Context){
	var users = []User{
		{ID:7,Name:"张三"},
		{ID:8,Name:"李四"},
		{ID:9,Name:"王五"},
	}
	c.JSON(200,users)
}

type User struct {
	ID int
	Name string
}
