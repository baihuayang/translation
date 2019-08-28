package main

import ("fmt"
"time")

func main(){
	ch := make(chan int, 10)
	go send(ch)
	go receive(ch)
	time.Sleep(1e9)
}

func send(ch chan int){
	for i:=0;i<10;i++{
		ch <- i*10
	}
}

func receive(ch chan int){
	for i:=0;i<10;i++{
		fmt.Println(<-ch)
	}
}