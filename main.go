package main

import (
	"runtime"
	"app/db"
	"app/server"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	db.Connect()
	server.Run()
}
