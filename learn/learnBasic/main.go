package main

import (
	"fmt"
	"reflect"
)

type (
	Myint int
)

type Myints []Myint

func (m Myints) String() {
}

type Xer interface {
	String()
}

func main() {
	var x = 0x111111
	fmt.Println((x))

	var m Myints = []Myint{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(m).Name()) // Kind: slice; Name: Myints

}
