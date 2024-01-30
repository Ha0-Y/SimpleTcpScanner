package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var addr string
	IP := "127.0.0.1"
	for port := 22; port < 1024; port++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			addr = fmt.Sprintf("%s:%d", IP, i)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Printf("%d close \n", i)
				return
			}
			fmt.Printf("%d open \n", i)
			err = conn.Close()
			if err != nil {
				panic(err)
			}
		}(port)
	}
	wg.Wait()
}
