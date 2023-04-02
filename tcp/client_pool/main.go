package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type connPool struct {
	idleConns map[int]net.Conn
	idleCount int
	lock      sync.Mutex
}

func (pool *connPool) getConn() net.Conn {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	if pool.idleCount == 0 {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			return nil
		}
		return conn
	}

	conn := pool.idleConns[pool.idleCount]
	delete(pool.idleConns, pool.idleCount)
	pool.idleCount--
	return conn
}

func (pool *connPool) putIdle(conn net.Conn) {
	pool.lock.Lock()
	defer pool.lock.Unlock()
	pool.idleCount++
	pool.idleConns[pool.idleCount] = conn
}

func main() {
	pool := connPool{
		idleConns: make(map[int]net.Conn),
		idleCount: 0,
		lock:      sync.Mutex{},
	}

	for {
		fmt.Fprintf(os.Stdout, "pool idle count: %d \n", pool.idleCount)
		conn := pool.getConn()
		n, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(os.Stdout, "written: %d \n", n)
		pool.putIdle(conn)
		time.Sleep(3 * time.Second)
	}
}
