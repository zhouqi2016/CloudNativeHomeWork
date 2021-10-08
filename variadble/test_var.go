package main

import "fmt"

func main() {
	test_other_func_const()
}

func var_claim() {
	//直接声明并赋值
	var a string = "hello, world"
	fmt.Println(a)
	//通过类型推断来赋值
	b := "hello, Tom"
	fmt.Println(b)
	//声明一个空值
	var c string
	fmt.Println(c)
	//对变量进行赋值
	c = "what is your name?"
	fmt.Println(c)
	//支持多变量声明
	var d, e int = 1, 2
	fmt.Printf("%v, %v", d, e)
}

func value_and_ref() {
	var a int = 1
	var b int = 1
	fmt.Println(&a)
	fmt.Println(&b)
	//可以看到他们在内存中的地址实际是不同的。这点和python有所不同，python做了相应的优化工作。
	fmt.Println(&a == &b)
}

func exchange_var() {
	//交换变量的快速使用方法和其他的语言使用temp临时变量不同
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("%v, %v", a, b)

}

//外部定义一个常量可以被访问到
const PI = 3.14

//这种是联合定义常量的方式
const (
	LEN1  = 55
	WIDHT = 66
)

func constant_var() {
	const PI = 3.14

}

func test_other_func_const() {
	fmt.Println(PI)
	//这个地方需要注意，即使常量定义在了后面，依然可以访问到。
	fmt.Println(LEN)
	fmt.Println(LEN1 * WIDHT)
}

const LEN = 55

//iota
