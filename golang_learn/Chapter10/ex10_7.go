package main

import (
"fmt"
//"strconv"
)

type stack struct{
	idx int
	nums [4]int
}

func (this *stack) pop() int{
	this.idx -= 1
	return this.nums[this.idx+1]
}

func (this *stack) push(num int){
	this.nums[this.idx] = num
	this.idx += 1
} 

func main(){
	s := stack{3, [4]int{1,2,3,4}}
	s.pop()
	s2 := new(stack)
	fmt.Println(s)
	fmt.Println(s2)
}

func (this *stack) String() string{
	return "11"
}



