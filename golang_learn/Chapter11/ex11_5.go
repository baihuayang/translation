package main

import(
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)
	go sendString(ch)
	go receiveString(ch)
}

func sendString(ch chan){
	ch <- "AAAAA"
	ch <- "BBBBB"
	ch <- "CCCCC"
	ch <- "DDDDD"
	ch <- "EEEEE"
}

func receiveString(ch chan){
	for{
		input := <- ch
		fmt.Printf("%s ", input)
	}
}