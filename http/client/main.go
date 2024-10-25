package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	/* 示例1：GET */
	{
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")

		/* 发请求收应答 */
		ack, err := http.Get("http://127.0.0.1:1280/")
		if err != nil {
			panic(err)
		}

		/* 读取应答正文 */
		ackBody, err := ioutil.ReadAll(ack.Body)
		/* 关闭应答正文，释放资源，无论是否异常 */
		ack.Body.Close()
		if err != nil {
			panic(err)
		}

		/* 输出应答状态 */
		fmt.Printf("HTTP Response StatusCode: %d\n", ack.StatusCode)
		fmt.Printf("HTTP Response Status: %s\n", ack.Status)

		/* 输出应答头域 */
		fmt.Printf("HTTP Response HEADER: %s\n", ack.Header.Get("my-http-head"))

		/* 输出应答正文 */
		fmt.Printf("HTTP Response BODY: %s\n", ackBody)
	}

	/* 示例2：POST */
	{
		fmt.Println("---------------------------------------------")

		/* 构建请求正文 */
		reqBody := strings.NewReader(`
			{
				"name": "test1280"
			}
		`)

		/* 创建请求对象 */
		req, err := http.NewRequest("POST", "http://127.0.0.1:1281/", reqBody)
		if err != nil {
			panic(err)
		}

		/* 设置请求头域 */
		req.Header.Set("Content-Type", "application/json")

		/* 发请求收应答 */
		ack, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}

		/* 读取应答正文 */
		ackBody, err := ioutil.ReadAll(ack.Body)
		/* 关闭应答正文，释放资源，无论是否异常 */
		ack.Body.Close()
		if err != nil {
			panic(err)
		}

		/* 输出应答状态 */
		fmt.Printf("HTTP Response StatusCode: %d\n", ack.StatusCode)
		fmt.Printf("HTTP Response Status: %s\n", ack.Status)

		/* 输出应答头域 */
		fmt.Printf("HTTP Response HEADER: %s\n", ack.Header.Get("my-http-head"))

		/* 输出应答正文 */
		fmt.Printf("HTTP Response BODY: %s\n", ackBody)
	}
}
