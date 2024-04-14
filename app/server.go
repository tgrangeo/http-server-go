package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()
	connection, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	b := make([]byte,1024)
	_,err = connection.Read(b)
	if err != nil{
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	request := string(b)
	status := strings.Split(request, "\r\n")
	path := strings.Split(status[0], " ")[1]
	if path == "/" {
		connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		connection.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
