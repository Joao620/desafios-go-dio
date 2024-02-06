package main

import (
	"fmt"
	"time"
)

func pong(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string)
	go pong(c)
	for {
		fmt.Println("ping")
		c <- "pong"
		time.Sleep(time.Second * 1)
	}
}
