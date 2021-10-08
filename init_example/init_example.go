package main

import (
    "fmt"
	"flag"
)

const (
    HELLO = "hello"
)

var (
    kubeconfig string
)


func main() {
    flag.Parse()
    b := kubeconfig
	fmt.Println(b)
}

func init() {
    flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
}
