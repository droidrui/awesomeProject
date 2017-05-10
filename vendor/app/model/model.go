package model

type Response struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

type Memo struct {
	Id       int `db:"id" json:"id"`
	Content  string `db:"content" json:"content"`
	LastTime string `db:"lastTime" json:"lastTime"`
}
