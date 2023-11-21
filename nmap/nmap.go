package main

import (
	"fmt"
	"log"

	"github.com/Ullaakut/nmap"
)

func main() {
	// 创建一个新的nmap扫描器实例
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("10.239.200.1-254"),
		nmap.WithPingScan(), // 添加ping扫描选项（等同于"-sn"）
	)
	if err != nil {
		log.Fatalf("无法创建nmap扫描器：%v", err)
	}

	// 运行扫描
	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("扫描失败：%v", err)
	}

	// 打印结果
	r := make([]string, 16)
	// fmt.Printf("主机列表：\n")
	// fmt.Printf("result: %v\n", result)
	for _, host := range result.Hosts {
		address := host.Addresses[0].String()
		r = append(r, address)
		// fmt.Println(host.Addresses[0])
	}
	fmt.Printf("r: %v\n", r)
}
