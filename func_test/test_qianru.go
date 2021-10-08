package main

import "fmt"


type notifier interface{
    notify()
}

type boy struct{
	*person
	level string
}

type person struct{
	name string
	age int
	email string
}

func (p person) notify(){
	fmt.Printf("my name emil is %s", p.email)
}

//这里如果实现了notify，那么就相当于实现了内部类的方法。无法实现内部类提升
//这个东西本质上就是实现了继承的概念。
func (b boy) notify(){
	fmt.Printf("my email is xxx %s", b.email)
}

func main(){
	//这里如果使用&boy那么b通过类型推断会判断出是一个指针，指针的调用可以和非指针类型一样。
	b := boy{
		person: &person{
			name: "tom",
			age: 19,
			email: "tom@gmail.com",
		},
		level: "super",
	}
	//这个地方其实代表内部类提升, 内部类提升其实从另外一个角度实现了继承。但又不是继承。
	b.notify()
}