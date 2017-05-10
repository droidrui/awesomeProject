package controller

import (
	"app/server"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func init() {
	server.Router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
	})
}
