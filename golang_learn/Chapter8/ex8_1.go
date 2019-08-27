package main

import "fmt"

func main(){
	map1 := make(map[int]string)
	map1[1]="Monday"
	map1[2]="Tuesday"
	map1[3]="Wondesday"
	map1[4]="Thursday"
	map1[5]="Friday"
	map1[6]="Saturday"
	map1[7]="Sunday"
	
	for _,value := range map1{
		if value == "Tuesday" || value == "Hollyday"{
			fmt.Println("yes")
		}
	}
}