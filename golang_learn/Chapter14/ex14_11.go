package main

import(
"fmt"
"time"
)

func main(){

}

func pump1(ch chan int){
	for i:=0;;i++{
		ch <- i*2
	}
}

func pump2(ch chan int){
	for i:=0;;i++{
		ch <- i+5
	}
}

func suck(ch1, ch2 chan int){
	for{
		select{
			case v := <-ch1:
				fmt.Printf("Received on channel 1: %d\n", v)
			case v := <-ch2:
				fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}