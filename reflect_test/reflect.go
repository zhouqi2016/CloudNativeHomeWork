package main

import "fmt"
import "reflect"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("hello, world")
	a := 100
	b := 200
	valuea := reflect.ValueOf(a)
	fmt.Println("type:", reflect.TypeOf(b))
	// 注意这里的valuea实际是一个reflect value类型
	fmt.Println("value:", valuea)
	value := valuea.Interface().(int)
	fmt.Println("real_value:", value)
	p1 := &Person{"name": "Tom", "age": 18}

}

func test_set_value(a interface{}) {
	value = reflect.ValueOf(a)
	fmt.Println("value:", value)
}
