package main

import (
	"fmt"
	"os"
    "github.com/TimotheeHrl/goservices/user"
)

func main() {

	arg := os.Args[1]
	fmt.Println(arg)
	GetUser(arg)
}
