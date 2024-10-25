package main

import (
	"embed"
	_ "embed"
	"log"
	"net/http"
)

//go:embed version.txt
var version string

//go:embed *.txt
var content embed.FS

func main() {
	//fmt.Printf("version: %q\n", version)

	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
