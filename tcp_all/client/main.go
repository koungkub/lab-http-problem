package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for {
		n, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(os.Stdout, "written: %d \n", n)
		time.Sleep(3 * time.Second)
	}
}
