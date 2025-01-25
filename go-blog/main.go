package main

import (
	"fmt"

	"github.com/khalidkhnz/sass/go-blog/config"
)

func main() {
	config.InitEnv()
	fmt.Println("hello blog")
}