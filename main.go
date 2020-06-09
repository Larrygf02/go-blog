package main

import (
	"fmt"

	"github.com/larrygf02/go-blog/controllers"
)

var server = controllers.Server{}

func main() {
	server.Initialize("postgres", "postgres", "123", "5432", "localhost", "bloggo")
	server.Run(":5000")
	fmt.Println("Hello world")
}
