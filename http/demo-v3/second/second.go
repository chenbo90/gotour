package main

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=========== i am second ==========")
		if len(r.Header) > 0 {
			for k, v := range r.Header {
				fmt.Printf("%s=%s", k, v[0])
			}
			dataType, _ := json.Marshal(r.Header)
			dataString := string(dataType)
			w.Write([]byte(dataString))
		}
	})
	fmt.Printf("start listen 0.0.0.0:8082\n")

	http.ListenAndServe("0.0.0.0:8082", nil)
}
