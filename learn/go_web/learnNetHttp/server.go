package learnnethttp

import "net/http"

// 如何创建一个http server？

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// r: 获取请求信息
	// w: 返回给客户端的内容
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("X-My-Header", "I am setting a header")
	w.Write([]byte("Hello World\n"))
}

// func main() {
// 	http.HandleFunc("/", HelloWorld)
// 	// 对于每一个请求，都会创建一个协程去处理
// 	http.ListenAndServe(":8080", nil)
// }
