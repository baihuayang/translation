package main

import "fmt"

func main(){
	ch := make(chan int)
	go tel(ch)
	for{
		select{
			case v:= <- ch:
				fmt.Printf("Received on channel 1: %d\n", v)
		}
	}
	
}

func tel(ch chan int){
	for i:=0;;i++{
		ch <- i
	}
}