package main

import "fmt"

func main() {
	a, b := 100, 200
	test_func(&a, &b)
	fmt.Println(a, b)
}

//交换两个值, 这里使用的是引用传递，对元素的操作会作用到真实的数据。
func test_func(a *int, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}
