package main

import (
	"errors"
	"fmt"
	"math"
)

func main(){
    ret1, err1 := sqrt(2)
	if err1!=nil {
		fmt.Print("errorrrrrr")
	}else{
		fmt.Print(ret1)
	}
	
}

func sqrt(f float64) (float64, error){
    if f<0 {
	    return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	
	return math.Sqrt(f), nil
}