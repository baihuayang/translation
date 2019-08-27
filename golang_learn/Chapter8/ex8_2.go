package main

import ("fmt" 
"sort")

func main(){
	drinkMap := map[string]int{"cola":3,"spirit":2,"fenda":4,"qixi":1}
	ss := make([]string, len(drinkMap))
	i := 0
	for key,_ := range drinkMap{
		ss[i] = key
		i++
	}
	sort.Strings(ss)
	for _,key := range ss {
		//fmt.Println(key)
		fmt.Println(drinkMap[key])
	}
	
}