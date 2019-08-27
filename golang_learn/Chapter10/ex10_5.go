package main

import "fmt"

type Base struct{
	id int
}

func (this *Base) Id() int{
	return this.id
}

func (this *Base) SetId(id int){
	this.id = id
}

type Person struct{
	Base
	FirstName string
	LastName string
}

type Employee struct{
	Person
	salary string
}

func main(){
	e := &Employee{Person{Base{123},"tom","yang"},"60000"}
	fmt.Println(e.Id())
}