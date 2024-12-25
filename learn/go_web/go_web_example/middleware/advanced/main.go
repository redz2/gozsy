package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 中间件，本身是一个函数类型，包裹handler处理函数
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 中间件工厂函数，用来生成中间件
// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// 使用中间件来增强handler功能
// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f // 经过各种中间件，返回一个新的HandlerFunc
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

// 有空就看看整个逻辑
func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}

// 总结
// 1. 中间件
// 增强函数功能（装饰器）

// 2. Pipeline（更加优雅的中间件写法）
// 使用中间件可以增强函数功能，不过使用中间件时会形成函数嵌套
// 通过Pipeline，也就是上面的Chain函数，使用时嵌套结构就变成扁平结构了
