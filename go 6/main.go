package main

import (
	"fmt"
	"time"
)

func main() {
	mychannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			mychannel <- i
			fmt.Println("Sent data:", i)
			time.Sleep(1 * time.Second)
		}
		close(mychannel)
	}()

	for {
		data, isOpen := <-mychannel
		if isOpen == false {
			break
		}
		fmt.Println("Received data:", data)
	}
}
