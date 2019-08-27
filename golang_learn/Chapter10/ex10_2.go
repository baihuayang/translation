package main

import("fmt")

type Rectangle struct{
	width int
	height int
	
}

func main(){
	m := Rectangle{3,5}
	fmt.Println(m.Area())
}

func (this *Rectangle) Area() int{
	return this.width*this.height
}