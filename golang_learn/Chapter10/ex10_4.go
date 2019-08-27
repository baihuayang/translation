package main

import("fmt")

type Car struct{
	wheelCount string
}

func (this *Car) numberOfWheels() string{
	return this.wheelCount
}

type Mercedes struct{
	Car
}

func main(){
	m := &Mercedes{Car{"4"}}
	fmt.Println(m.numberOfWheels())
}