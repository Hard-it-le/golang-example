package main

import (
	"fmt"
	. "net/http"
)

func home(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w ResponseWriter, r *Request) {
	fmt.Fprintf(w, "这是订单")
}

func main() {
	HandleFunc("/", home)
	HandleFunc("/user", user)
	HandleFunc("/user/create", createUser)
	HandleFunc("/order", order)
	ListenAndServe(":8080", nil)
}

type Server interface {
	Route(pattern string, handlerFunc HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}
