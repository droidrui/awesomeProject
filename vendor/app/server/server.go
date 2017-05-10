package server

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"app/mw"
	"app/controller"
)

var Router *httprouter.Router

func Run() {
	fmt.Println(time.Now(), "Running HTTP", httpAddress())

	Router = httprouter.New()

	log.Fatal(http.ListenAndServe(httpAddress(), handleMW(Router)))

	controller.Load()
}

func httpAddress() string {
	return ":80"
}

func handleMW(h http.Handler) http.Handler {
	h = mw.LogRequest(h)
	return h
}
