package main

import(
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)

	go sendString(ch)
	go receiveString(ch)
	time.Sleep(1e9)
}

func sendString(ch chan string){
	ch <- "AAAAA"
	ch <- "BBBBB"
	ch <- "CCCCC"
	ch <- "DDDDD"
	ch <- "EEEEE"
}

func receiveString(ch chan string){
	for{
		input := <- ch
		fmt.Printf("%s ", input)
	}
}