package main

import "fmt"

func main(){
	ch := make(chan int)
	go sum(ch,1,2)
	fmt.Println(<-ch)
}

func sum(ch chan int, x int, y int){
	ch <- x+y
}