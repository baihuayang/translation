package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}

var content	string
var vc VCard
func main() {
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	
	inReader := bufio.NewReader(file)
	enc := gob.NewDecoder(inReader)
	err := enc.Decode(&vc)
	if err != nil {
		log.Println(err)
		log.Println("Error in encoding gob")
	}
	fmt.Println(vc)
}