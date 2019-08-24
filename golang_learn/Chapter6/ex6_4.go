package main

import(
	"fmt"
)

func main(){
   ppp(10)
}

func ppp(n int){
    if n>0 {
	   fmt.Println(n)
	   ppp(n-1)
	}
}