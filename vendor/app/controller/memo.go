package controller

import (
	"app/server"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"app/db"
	"app/errcode"
	"app/model"
	"fmt"
)

func init() {
	server.Router.GET("/memo", getMemo())

	server.Router.POST("/memo", createMemo())

	server.Router.DELETE("/memo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	})

	server.Router.PUT("/memo", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	})
}

func createMemo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		content := r.FormValue("content")
		if content == "" {
			server.SendError("content is nil", w, errcode.PARAM_INVALID)
			return
		}
		result, err := db.SQL.Exec("INSERT INTO memo(content) VALUE(?)", content)
		if err != nil {
			server.SendError(err, w, errcode.SERVER_ERROR)
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			server.SendError(err, w, errcode.SERVER_ERROR)
			return
		}
		memo := model.Memo{}
		err = db.SQL.Get(&memo, "SELECT * FROM memo WHERE id=?", fmt.Sprint(id))
		if err != nil {
			server.SendError(err, w, errcode.SERVER_ERROR)
			return
		}
		server.SendSuccess(w, &memo)
	}
}

func getMemo() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	}
}
