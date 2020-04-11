package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go count("sheep", c)

	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func count(t string, c chan string) {

	for i := 1; i <= 5; i++ {
		c <- t
		time.Sleep(time.Second * 1)
	}

	close(c)

}
