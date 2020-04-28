package main

import (
	"fmt"

	"github.com/larrygf02/go-blog/config"
)

func main() {
	fmt.Println("Hello world")
	config.InitialMigration()
}
