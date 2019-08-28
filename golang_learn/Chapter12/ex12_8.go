package main

import(
"fmt"
"os"
"strings"
)

func main(){
	str := "Hello "
	if len(os.Args)>1{
		str += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(str)
}