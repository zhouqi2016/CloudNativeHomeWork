package main

import "fmt"

func main() {
	fmt.Println(b)
	fmt.Println(a, b, c, d, e)
}

const (
	z = "ha"
	f
)

//注意下面如果被终端了，紧接着的地方需要重新写一个iota否则后面都是cccc
const (
	a = 0 + iota
	b
	c = "cccc"
	d = iota
	e
)
