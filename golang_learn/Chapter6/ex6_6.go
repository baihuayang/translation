package main

import "fmt"

func main(){
  
	p2 := fibonaqi()
	for i:=0;i<10;i++ {
		fmt.Println(p2())
		
	}
}


func fibonaqi() func() int{
	var x int
	var y int
	var t int
	y = 1
	return func() int{
			
		   t = x	
	       x = y
		   y = t + y
		   return y
	}
}