package main

import("fmt")

func main(){
	var fv int
	func(){
		fmt.Print("hello world")
		fmt.Printf(" - fv is of type %T ", fv)
	}()
}