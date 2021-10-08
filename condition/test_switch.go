package main

import "fmt"

func main() {
	var grade int = 100
	var level string
	//注意这里如果没有表达式，则默认为true，case后面的表达式会和true做比较。
	//注意fallthrough表达式，其他语言如果不使用break就会一直往下执行，一直匹配。这里相反。
	//如果写fallthrough则会继续往下匹配，这个是符合实际编程情况的。
	switch {
	case grade >= 90 && grade <= 100:
		level = "优秀"
		//fallthrough是写在这里的
	case grade < 90 && grade >= 80:
		level = "良好"
	case grade < 90 && grade >= 70:
		level = "及格"
	default:
		level = "不及格"
	}
	fmt.Println(level)

}
