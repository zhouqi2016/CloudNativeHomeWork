package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	Readable = 1 + iota
	Writeable
	Pending
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	fmt.Println("hello, world!")
	fmt.Println(runtime.GOOS)
	fmt.Println(Writeable)
	fmt.Println(Pending)
}
