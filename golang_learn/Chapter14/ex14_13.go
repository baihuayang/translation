package main

import (
	"fmt"
	"os"
)

func tel(ch1 chan int, ch2 chan bool) {
	for i := 0; i < 15; i++ {
		ch1 <- i
	}
	ch2 <- true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go tel(ch1,ch2)
	for {
		select{
			case v:= <- ch1:
				fmt.Printf("Received on channel 1: %d\n", v)
			case <- ch2:
				os.Exit(0)
		}
	}
}