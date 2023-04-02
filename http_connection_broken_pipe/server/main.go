package main

import (
	"log"
	"net"
	"os"
)

func main() {
	server()
}

func server() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("server", err)
		os.Exit(1)
	}
	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Fatal("server", err)
	}

	conn.Close()
}
