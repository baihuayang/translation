package main

import(
	"fmt"
)

func main(){
    var a int
    var b int
	a=1
	b=2
	var c int
	var d int
	var e int
	c,d,e = test(a,b)
	fmt.Printf("%d%d%d",c,d,e)
}

func test (a int, b int) (int, int, int){
	return a+b, a-b, a*b
}