package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new connection")
			buf := make([]byte, 5)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					if err.Error() == "EOF" {
						break
					}
					fmt.Println(err)
				}
				if n == 0 {
					continue
				}
				fmt.Fprintf(os.Stdout, "readed size: %d, %s \n", n, buf)
			}
		}(conn)
	}
}
