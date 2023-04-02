package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	ctx := context.Background()
	trace := httptrace.ClientTrace{
		PutIdleConn: func(err error) {
			fmt.Println("put idle conn")
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("connect start")
		},
		GetConn: func(hostPort string) {
			fmt.Println("get conn")
		},
	}
	req := request.WithContext(httptrace.WithClientTrace(ctx, &trace))
	for {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		io.ReadAll(resp.Body)
		time.Sleep(3 * time.Second)
	}
}
