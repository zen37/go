package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 3)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}
