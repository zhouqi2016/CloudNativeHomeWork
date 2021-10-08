package main

import "fmt"

func main() {
	var grade int = 100
	var level string
	//这里注意，else if 必须在 } 后，而且格式固定。
	if grade == 100 {
		level = "优秀"
	} else if grade < 100 || grade >= 90 {
		level = "良好"
	} else {
		level = "及格"
	}
	fmt.Println(level)
	test_say_hello()
}

func test_say_hello() {
	//这种方式其实是golang比较常用的一种表达式，比较明了，和python的try catch不同。
	if err, str := say_hello(); err == nil {
		fmt.Println(str)
	} else if err != nil {
		//这里主要用来处理异常的情况
		fmt.Println("He do not say hello")
	}
}

func say_hello() (error, string) {
	return nil, "hello, world"
}
