package controller

import (
	"app/server"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func init() {
	server.Router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Welcome!")
	})
}
