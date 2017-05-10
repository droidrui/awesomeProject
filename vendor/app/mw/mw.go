package mw

import (
	"net/http"
	"fmt"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Host, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
