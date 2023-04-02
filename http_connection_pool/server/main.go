package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello, world"))
}

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: handler{},
		ConnState: func(conn net.Conn, state http.ConnState) {
			fmt.Println(conn.LocalAddr(), conn.RemoteAddr(), state.String())
		},
	}

	log.Fatal(srv.ListenAndServe())
}
