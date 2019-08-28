package main

import (
	"fmt"
)


func main() {
	out := make(chan int)
	out <- 2
	fmt.Println("xx")
	//go f1(out)
}