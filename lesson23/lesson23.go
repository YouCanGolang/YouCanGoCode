package main

import (
	"io"
	"net/http"
)

var (
	userMap map[string]string
)

func init() {
	userMap = make(map[string]string)
}

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.ListenAndServe("127.0.0.1:8088", nil)
}

func register(rsp http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	_, ok := userMap[username]
	if ok {
		io.WriteString(rsp, "已经注册过了")
		return
	}

	userMap[username] = password
	io.WriteString(rsp, "success")
}

func login(rsp http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	cachePassword, ok := userMap[username]
	if !ok {
		io.WriteString(rsp, "用户名未注册")
		return
	}

	if cachePassword != password {
		io.WriteString(rsp, "密码错误")
		return
	}
	io.WriteString(rsp, "success")
}
