package controller

import (
	"app/server"
	"net/http"
	"github.com/julienschmidt/httprouter"
	_"fmt"
	"html/template"
	"app/errcode"
)

func init() {
	server.Router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		//fmt.Fprintln(w, "Welcome!")
		if err := template.Must(template.ParseFiles("index.html")).Execute(w, nil); err != nil {
			server.SendError(err, w, errcode.SERVER_ERROR)
		}
	})
}
