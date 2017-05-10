package server

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"app/mw"
)

var Router *httprouter.Router

func init() {
	Router = httprouter.New()
}

func Run() {
	fmt.Println(time.Now(), "Running HTTP", httpAddress())

	log.Fatal(http.ListenAndServe(httpAddress(), handleMW(Router)))
}

func httpAddress() string {
	return ":80"
}

func handleMW(h http.Handler) http.Handler {
	h = mw.LogRequest(h)
	return h
}
