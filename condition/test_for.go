package main

import "fmt"

func main() {
	//这里可以用作while使用
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	var array = []string{"a", "b", "c"}
	//这里用来遍历数组
	for _, value := range array {
		fmt.Println(value)
	}

}
