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
	//m := new(Simple)
	m := Simple{1}
	m.id = 1
	var s Simpler
	s = m
	fmt.Println(s)
	 
}