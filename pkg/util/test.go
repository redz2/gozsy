package util

import (
	"fmt"
	"io"
	"net/http"
)

const (
	VictoriaUrl string = "sdad"
)

// 首字母必须大写!!!!!!
func SayHello() {
	fmt.Println("sasdasd")
}

func HttpGet(url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := io.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
	return
}
