package main

import (
	"fmt"
)

func main(){
	server := NewApiServer(":3001");
	server.Run()

	fmt.Println("Yeah Buddy!")
}