package main

import (
	"fmt"

	"github.com/khalidkhnz/sass/go-ecom/config"
)

func main() {
	config.InitEnv()
	fmt.Println("hello ecom")
}