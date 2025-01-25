package main

import (
	"fmt"

	"github.com/khalidkhnz/sass/go-sass/config"
)

func main() {
	config.InitEnv()
	fmt.Println("hello sass")
}