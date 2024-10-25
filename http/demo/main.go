package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/chat/completions", func(w http.ResponseWriter, r *http.Request) {
		// 允许读取请求体
		if r.Method == "POST" {
			// 读取请求体
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "无法读取请求体", http.StatusBadRequest)
				return
			}

			// 解析请求体（假设是JSON格式）
			// 这里需要根据实际的请求体格式进行解析和处理
			// 例如，如果请求体是JSON，可以使用json.Unmarshal来解析

			// 假设解析后的数据存储在data变量中
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			fmt.Printf("body:%v\n", data)

			if err != nil {
				http.Error(w, "无法解析请求体", http.StatusBadRequest)
				return
			}

			// 添加一个新的参数到请求体中
			// 假设我们要添加的参数是"newParam": "newValue"
			data["newParam"] = "newValue"

			// 将更新后的数据转换回请求体
			updatedBody, err := json.Marshal(data)
			if err != nil {
				http.Error(w, "无法序列化请求体", http.StatusInternalServerError)
				return
			}

			// 创建一个新的请求
			// 注意：这里需要替换为实际的目标URL
			targetURL := "https://api.moonshot.cn/v1/chat/completions"
			req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(updatedBody))
			if err != nil {
				http.Error(w, "无法创建新请求", http.StatusInternalServerError)
				return
			}

			// 复制原始请求的Header到新请求
			for key, values := range r.Header {
				for _, value := range values {
					fmt.Println("header:" + value)
					req.Header.Add(key, value)
				}
			}

			// 发送请求到另一个接口
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				http.Error(w, "请求失败", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			// 读取响应体
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "无法读取响应体", http.StatusInternalServerError)
				return
			}

			// 将响应写回原始请求的响应中
			w.Write(respBody)
		} else {
			http.Error(w, "不支持的方法", http.StatusMethodNotAllowed)
		}
	})
	fmt.Printf("start listen 0.0.0.0:8080\n")

	http.ListenAndServe("0.0.0.0:8080", nil)
}

//func main() {
//	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("打印Header参数列表：")
//		if len(r.Header) > 0 {
//			/*      for k,v := range r.Header {
//			        fmt.Printf("%s=%s", k, v[0])
//			}*/
//			dataType, _ := json.Marshal(r.Header)
//			dataString := string(dataType)
//			w.Write([]byte(dataString))
//		}
//	})
//	fmt.Printf("start listen 0.0.0.0:8080\n")
//
//	http.ListenAndServe("0.0.0.0:8080", nil)
//}
