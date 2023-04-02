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

	// write to make the connection closed on the server side
	if _, err := conn.Write([]byte("a")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate an RST packet
	if _, err := conn.Write([]byte("b")); err != nil {
		log.Printf("client: %v", err)
	}

	time.Sleep(1 * time.Second)

	// write to generate the broken pipe error
	if _, err := conn.Write([]byte("c")); err != nil {
		log.Printf("client: %v", err)
		if errors.Is(err, syscall.EPIPE) {
			log.Print("This is broken pipe error")
		}
	}
}
