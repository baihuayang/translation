package main

import "fmt"

func main(){
	s := make([]int,100)
	for i,_ :=range s{
     	s[i] = i
	}
	bo := func(x int) bool{
		return x%2==0
	}
	fmt.Println("xxxxxxxxxxxxxxxxx")
	
	t := Filter(s, bo)
	for _,v :=range t{
		fmt.Println(v)
	}
	
}

func Filter(x []int,ff func(int) bool) []int{
	s := make([]int, 0) 
	for _,v := range x{
		if ff(v){
			s = append(s,v)
		}
	}
	return s;
}