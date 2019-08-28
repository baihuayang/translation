package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string,10)
	go send(ch)
	go receive(ch)
	time.Sleep(2e10)
}

func send(ch chan string){
	time.Sleep(1e10)
	ch <- "hahaha"
	fmt.Println("sended")
}

func receive(ch chan string){
	str := <- ch
	fmt.Println("received " + str)
}