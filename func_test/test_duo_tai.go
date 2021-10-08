package func_test

import "fmt"

type  Person interface {
	SayHello()
}

type Boy struct {
	name string
	age int
}

func (boy *Boy)SayHello() {
    fmt.Printf("hello, my name is %s", boy.name)
}

type Girl struct{
	name string
	age int
}

func (g *Girl)SayHello() {
    fmt.Printf("hello, i am a girl, my name is %s", g.name)
}

func say_hello(p Person){
    p.SayHello()
}

func main(){
	b := &Boy{name: "jackson", age: 18}
	say_hello(b)
	g := &Girl{name: "lili", age: 20}
    say_hello(g)
    var tt Person
	//golang是支持多态的，这个很重要，多态的实现。
	tt = &Boy{name: "jackson", age: 18}
	tt.SayHello()
}

