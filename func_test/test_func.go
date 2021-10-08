package test_func

import "fmt"

func main() {
    boy := &Boy{name: "tom", age: 18}
    boy.SetName("jackson")
    boy.SayHello()
}

type  Person interface {
    SetName(new_name string)
	SayHello()
}

type Boy struct {
    name string
	age int
}
//对于设置和修改类型的必须使用指针，否则操作的知识boy的副本，而不是真正的值。
func (p *Boy) SetName(new_name string){
    p.name = new_name
}

func (p *Boy) SayHello(){
    fmt.Printf("hello, world! my name is %s \n", p.name)
}
