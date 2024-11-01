// package main

// import (
// 	"github.com/mohuishou/go-test-multi-module/a"
// )

// // var CLI struct {
// // 	Rm struct {
// // 		Force     bool `help:"Force removal."`
// // 		Recursive bool `help:"Recursively remove files."`

// // 		Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
// // 	} `cmd:"" help:"Remove files."`

// // 	Ls struct {
// // 		Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
// // 	} `cmd:"" help:"List paths."`
// // }

// func main() {
// 	// ctx := kong.Parse(&CLI)
// 	// switch ctx.Command() {
// 	// case "rm <path>":
// 	// case "ls":
// 	// default:
// 	// 	panic(ctx.Command())
// 	// }
// 	// mymoduletest.PrintHelloWorld()
// 	a.A()

// }
package main

import "github.com/redz2/gozsy/learn/hello"

func main() {
	hello.PrintHelloWorld()
}
