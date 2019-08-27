package main

import "fmt"

type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

func Sort(data Sorter) {
	fmt.Println("Sort!")
	bbb := data.(*personArray)
    for pass := 1; pass < data.Len(); pass++ {
        for i := 0;i < data.Len() - pass; i++ {
			fmt.Print(bbb.data[i])
			fmt.Print(" --- ")
			fmt.Print(bbb.data[i+1])
            if data.Less(i+1, i) {
				fmt.Println("swap!")
                data.Swap(i, i + 1)
            } else{
				fmt.Println(" NO SWAP !")
			}
        }
    }
}

type Person struct{
	firstName string
	lastName string
}

type personArray struct {
	data []*Person
}

func (this personArray) Len() int{
	return len(this.data)
}

func (this personArray) Less(i, j int) bool{
	return this.data[i].firstName < this.data[j].firstName
}

func (this personArray) Swap(i, j int){
	 this.data[i], this.data[j] = this.data[j],this.data[i] 
}

func main(){
	Tom1 := Person{"1tom1","fff1"}
	Tom2 := Person{"3tom3","fff3"}
	Tom3 := Person{"2tom2","fff2"}
	Tom4 := Person{"4tom4","fff4"}
	data := []*Person{&Tom1,&Tom2,&Tom3,&Tom4}
	a := personArray{data}
	Sort(&a)
	for _,v := range data{
		fmt.Println(v)
	}
	if "1tom1" < "3tom3" {
		fmt.Println("xx")
	} else {
		fmt.Println("dd")
	}
}