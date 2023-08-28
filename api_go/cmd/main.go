package main

import (
	"demo"
)

func main() {
	server := demo.NewServer()
	server.Start(":1323")
}
