package main

import (
	`fmt`
	`net`
)

func main() {
	var addr string
	var err error
	var conn net.Conn
	IP := "127.0.0.1"
	for port := 22; port < 1024; port++ {
		addr = fmt.Sprintf("%s:%d", IP, port)
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			fmt.Printf("%d close\n", port)
			continue
		}
		fmt.Printf("%d open\n", port)
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}
}
