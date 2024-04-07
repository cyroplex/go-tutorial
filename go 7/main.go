package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		data1 := <-messages
		fmt.Println("First:", data1)
		data2 := <-messages
		fmt.Println("Second:", data2)
		data3 := <-messages
		fmt.Println("Third:", data3)
	}()

	messages <- "hello"
	messages <- "world"
	messages <- "world2"

	time.Sleep(1 * time.Second)
}
