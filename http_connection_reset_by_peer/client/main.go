package main

import (
	"errors"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	client()
}

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("client", err)
	}

	if _, err := conn.Write([]byte("ab")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second) // wait for close on the server side

	data := make([]byte, 1)
	if _, err := conn.Read(data); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.ECONNRESET) {
			log.Print("This is connection reset by peer error")
		}
	}
}
