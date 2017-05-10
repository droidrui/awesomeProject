package server

import (
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"app/mw"
	"app/model"
	"encoding/json"
	"app/errcode"
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

func SendError(msg interface{}, w http.ResponseWriter, code int) {
	log.Println(msg)
	resp := &model.Response{
		Code: code,
		Msg:  errcode.ErrMap[code],
		Time: fmt.Sprint(time.Now().Unix()),
	}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "JSON Marshal Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(json))
	w.Write(json)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	resp := &model.Response{
		Code: errcode.SUCCESS,
		Msg:  errcode.ErrMap[errcode.SUCCESS],
		Time: fmt.Sprint(time.Now().Unix()),
		Data: data,
	}
	json, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "JSON Marshal Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(json))
	w.Write(json)
}
