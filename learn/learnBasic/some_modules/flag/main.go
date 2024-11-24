package main

import (
	"flag"
	"fmt"
)

func main() {
	// var name = flag.String("name", "everyone", "The greeting object.") // 返回*string
	// flag.Parse()
	// fmt.Printf("Hello, %s\n", *name)

	// name := flag.String("name", "everyone", "The greeting object.")

	var name string
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %s\n", name)
}
