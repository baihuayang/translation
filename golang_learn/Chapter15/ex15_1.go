package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var m map[string]int

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	m = make(map[string]int)
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		if strings.Contains(string(buf),"SH"){
			os.Exit(0)
		}
		if strings.Contains(string(buf), ": WHO") {
			DisplayList()
		}
		input := string(buf)
		clName := input[0 : ix-1]
		m[string(buf)]=1
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range m {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}