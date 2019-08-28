package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string, 10)
	go send(ch)
	go receive(ch)
	time.Sleep(3e9)
}

func send(ch chan string){
	for i:=1;i<=10;i++{
		ch <- "111"
		fmt.Println("ssss")
	}
}

func receive(ch chan string){
	
		str := <-ch
		fmt.Println(str)
	
}