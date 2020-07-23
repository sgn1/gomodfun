package main

import (
	"github.com/davecgh/go-spew/spew"
)

func Add(a, b int) int {
	spew.Dump(a, b)
	return a + b
}

func Mult(a, b int) int {
	return a * b
}

func main() {
	spew.Dump(Add(1, 2))
}
