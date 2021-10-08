package main

import "fmt"
import "sync"
import "runtime"

func main(){
	//修改调度器可以使用的逻辑处理器的数量。
    runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	//这里的值为1，说明需要等待1个goroutine结束。
	wg.Add(2)

    go func(){
		defer wg.Done()
        fmt.Println("sleep 1h")
		for i :=0; i<100; i++{
			fmt.Println(i)
		}
	}()

	go func(){
		defer wg.Done()
		for i:=100; i<200; i++{
			fmt.Println(i)
		}
		fmt.Println("sleep 2h")
	}()
	fmt.Println("Wait to close")
	wg.Wait()
}
