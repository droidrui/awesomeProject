package main

import (
	"runtime"
	"net/http"
	"fmt"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("sourse")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}