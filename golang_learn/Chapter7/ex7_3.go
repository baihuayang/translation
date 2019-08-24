package main

import "fmt"

func main(){
	
	Append()
}

func Append(slice, data []byte) [] byte{
	copy := slice
	for ;len(slice)+len(byte) > cap(slice);{
		slice = make([]byte, 2*cap(slice))
	}
	slice = copy
	slice(len(copy):) = byte
}