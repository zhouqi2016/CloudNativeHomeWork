package main

import (
    "os"
	"fmt"
)

func main() {
    fmt.Println(os.Getenv("DOCKER_DRIVER"))
}

