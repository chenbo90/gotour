package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	r := gin.Default()
	logrus.SetReportCaller(true)
	r.GET("ginusers", listUser)
	r.POST("create", createUser)
	r.POST("addPabComponentSet", createPabComponentSet)
	r.GET("deletePabComponentSet", deletePabComponentSet)
	r.Run("localhost:8888")
}

func listUser(c *gin.Context) {
	var users = []User{
		{ID: 7, Name: "张三"},
		{ID: 8, Name: "李四"},
		{ID: 9, Name: "王五"},
	}
	c.JSON(200, users)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	marshal, _ := json.Marshal(user)
	logrus.Infof("[user]入参:%s", marshal)
	c.JSON(200, user)

}

func createPabComponentSet(c *gin.Context) {
	var vo PabComponentSetVo
	if err := c.ShouldBindJSON(&vo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	marshal, _ := json.Marshal(vo)
	logrus.Infof("[PabComponentSetVo]入参:%s", marshal)
	// 构建Response结构体
	resp := Response{
		Code:    200,
		Data:    vo,
		Message: "sucess",
	}
	c.JSON(200, resp)

}

func deletePabComponentSet(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	logrus.Infof("[删除组件集的id]入参:%v", id)
	// 构建Response结构体
	resp := Response{
		Code:    200,
		Data:    id,
		Message: "sucess",
	}
	c.JSON(200, resp)

}

type User struct {
	ID   int
	Name string `json:"name",validate:"required`
}

type PabComponentSet struct {
	Name        string            `json:"name"`
	Code        string            `json:"code"`
	Type        string            `json:"type"`
	Comment     string            `json:"comment"`
	CreatedBy   string            `json:"createdBy"`
	CreatedTime int64             `json:"createdTime"`
	TenantCode  string            `json:"tenantCode"`
	Params      map[string]string `json:"params"`
}

type PabComponentSetVersion struct {
	Version     string            `json:"version"`
	Template    string            `json:"template"`
	CreatedBy   string            `json:"createdBy"`
	CreatedTime int64             `json:"createdTime"`
	TenantCode  string            `json:"tenantCode"`
	Params      map[string]string `json:"params"`
}

type PabComponentSetVo struct {
	PabComponentSet        PabComponentSet        `json:"pabComponentSet"`
	PabComponentSetVersion PabComponentSetVersion `json:"pabComponentSetVersion"`
}

// Response 是一个统一的返回结构体
type Response struct {
	Code    int         `json:"code"` // 状态码
	Data    interface{} `json:"data"` // 数据，可以是任意类型
	Message string      `json:"msg"`  // 消息或错误信息
}
