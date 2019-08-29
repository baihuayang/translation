package main

import (
	"net/http"
	"log"
	"net"
	"net/rpc"
	"time"
	"mypkg"
)

func main() {
	calc := new(mypkg.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}