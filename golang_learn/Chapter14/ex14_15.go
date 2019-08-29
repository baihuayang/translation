package main

import "fmt"

var out = make(chan int)

func main(){
	aaa()
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
	fmt.Println(generateInteger())
}

func aaa(){
	go func(){
		for i:=1;;i++{
			out <- i
		}
	}()
	
}

func generateInteger() int{
	return <-out
}