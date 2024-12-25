package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// httprouter
// 优点: 不需要关心结尾的斜杠

func HomeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "home\n")
}

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "login\n")
}

func main() {
	// 创建路由器
	router := httprouter.New()
	// 路由分配
	router.GET("/", HomeHandler)
	router.POST("/login", LoginHandler)

	http.ListenAndServe(":8080", router)
}
