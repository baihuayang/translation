package main
import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]int
	for i:=0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	
	arr2 = arr1
	arr2[2] = 100
	for i,v := range arr1 {
		fmt.Printf("Array at index %d is %d\n", i, v)
	}
	
	
}