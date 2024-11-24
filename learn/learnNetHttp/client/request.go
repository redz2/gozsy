package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func HttpRequest(method string, url string, body io.Reader) {

	// 1，创建一个http.Client实例
	client := &http.Client{
		Timeout: 100 * time.Millisecond,
	}

	// 2，创建一个http的请求
	payload := strings.NewReader(`{"data": 24}`)
	r, _ := http.NewRequest(method, url, payload)
	// 自定义header
	r.Header.Add("Content-Type", "application/json")
	// 自定义cookie
	// ...

	// 3，发送请求并接收响应
	if resp, err := client.Do(r); err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		defer resp.Body.Close()

		// 4，读取响应内容
		d, _ := io.ReadAll(resp.Body)
		fmt.Printf("string(d): %v\n", string(d))
	}

	// 5. 解析JSON
	// json只是字符串的一种特定格式，必须要解析成结构体
}

// 复用*http.Client会有性能瓶颈吗？
