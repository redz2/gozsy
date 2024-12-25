package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Rm struct {
		Force     bool `help:"Force removal."`
		Recursive bool `help:"Recursively remove files."`

		Other string   `arg:"" name:"other" help:"Other params" type:"other"`
		Paths []string `arg:"" name:"path" help:"Paths to remove." type:"path"`
	} `cmd:"" help:"Remove files."`

	Ls struct {
		Paths []string `arg:"" optional:"" name:"path" help:"Paths to list." type:"path"`
	} `cmd:"" help:"List paths."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "rm <other> <path>":
		fmt.Println("rm")
		fmt.Println(ctx.Command())
		fmt.Println(ctx.Args) // 参数数组
	case "ls", "ls <path>":
		fmt.Println("ls")
		fmt.Println(ctx.Command())
		fmt.Println(ctx.Args) // 参数数组
	default:
		fmt.Println("default")
		fmt.Println(ctx.Command())
		// panic(ctx.Command())
	}
}
