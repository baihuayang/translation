package main

import (
	"fmt"
	"net/http"
)

type Hello struct{
}


func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[len("/hello/"):])
}

func ShouthelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[len("/shouthello/"):])
}


func main(){
	http.HandleFunc("/hello/",HelloServer)
	http.HandleFunc("/shouthello/",ShouthelloServer)
	http.ListenAndServe("localhost:8080", nil) 
}