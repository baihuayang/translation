package main

import (
	"fmt"
	"net/http"
)

type Hello struct{
}

func (m *Hello)ServeHTTP(w http.ResponseWriter, req *http.Request){
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main(){
	hello := new(Hello)
	http.Handle("/",hello)
	http.ListenAndServe("localhost:8080", nil) 
}