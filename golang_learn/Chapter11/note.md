package main

import "fmt"

type Simpler interface{
	get() int
	set(int)
}

type Simple struct{
	id int
}

func(this *Simple) get() int{
	return this.id
}

func(this *Simple) set(id int){
	this.id = id
}

func main(){
	Simpler m = Simple{1}
	fmt.Println(m)
	 
}