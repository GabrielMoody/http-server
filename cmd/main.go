package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Printf("Connected client from %s\n", conn.RemoteAddr())

		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 0\r\n"))
	}

}
