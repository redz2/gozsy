package main

import (
	"fmt"
	"net/http"
)

// net/http自带的mux

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "API")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "你好,世界!")
	})

	server := &http.Server{Addr: ":8080"}
	server.Handler = mux
	server.ListenAndServe()
}
