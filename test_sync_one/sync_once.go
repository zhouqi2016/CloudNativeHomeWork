package main

import "fmt"
import "sync"

type Singlone struct{}

var once sync.Once

func getSinglone() {
	once.Do(func() {
		singlone := new(Singlone)
	})
	return singletone
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			singletone = getSingletone()
			fmt.Println(&singletone)
		}()
	}
}
