package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	defer wg.Done()

	/* 监听一个端口方式 */
	go func() {
		/* 使用http.DefaultServeMux */
		http.HandleFunc("/", handle)
		http.ListenAndServe(":1280", nil)
	}()

	/* 监听多个端口方式 */
	go func() {
		/* 自定义ServeMux */
		mux := http.NewServeMux()
		mux.HandleFunc("/", handle)
		http.ListenAndServe(":1281", mux)
		http.ListenAndServe(":1282", mux)
	}()

	go func() {
		/* 自定义ServeMux */
		mux := http.NewServeMux()
		mux.HandleFunc("/", handle)
		http.ListenAndServe(":1282", mux)
	}()

	/* 测试代码，保证主线程/协程不退出 */
	//time.Sleep(time.Second * 60 * 10)

	//wg.Wait()
}

func handle(w http.ResponseWriter, r *http.Request) {
	/* 读取发请求客户端地址 */
	fmt.Printf("Receive HTTP REQ FROM: %s\n", r.RemoteAddr)
	/* 读取请求方法 */
	fmt.Printf("Method: %s\n", r.Method)
	/* 读取请求URL */
	fmt.Printf("URL: %s\n", r.URL)

	/* 读取请求头域 */
	for k, vs := range r.Header {
		for _, v := range vs {
			fmt.Printf("HTTP HEADER: %s\t\t:%s\n", k, v)
		}
	}

	/* 读取请求正文 */
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	/* 输出请求正文 */
	fmt.Printf("HTTP REQ BODY:\n%s\n", reqBody)

	/* 设置HTTP应答头 */
	w.Header().Set("my-http-head", fmt.Sprintf("HTTP-REQ-METHOD-%s", r.Method))

	/* 生成应答正文 */
	ackBody, err := json.Marshal(r.Header)
	if err != nil {
		panic(err)
	}
	/* 设置应答状态 */
	w.WriteHeader(201)
	/* 设置应答正文 */
	w.Write(ackBody)

}
