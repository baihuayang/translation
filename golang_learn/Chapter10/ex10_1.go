package main

import("fmt")

type Address struct{
	address string
}

type VCard struct{
	name string
	birthDate string 
	address *Address
}

func main(){
	m := VCard{"tom","881111",&Address{"西土城路"}}
	fmt.Println(m.address.address)
}