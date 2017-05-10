package main

import (
	"runtime"
	"app/db"
	"app/server"
	"app/controller"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	db.Connect()
	controller.Load()
	server.Run()
}
