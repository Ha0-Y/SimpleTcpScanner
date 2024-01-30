package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"time"
)

// chan <- 运算会阻塞

var (
	hostname  string
	startPort int
	endPort   int
)

func worker(ports, result chan int) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", hostname, port)
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err != nil {
			result <- 0
			continue
		}
		err = conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v", err)
		}
		result <- port
	}
}

func main() {
	flag.StringVar(&hostname, "h", "", "hostname")
	flag.IntVar(&startPort, "start-port", 21, "scanning start port")
	flag.IntVar(&endPort, "end-port", 65535, "scanning end port")
	flag.Parse()
	//hostname = "127.0.0.1"
	//startPort = 21
	//endPort = 1024

	fmt.Println("Start scanning...")
	ports := make(chan int, 100)
	res := make(chan int)
	var openPorts []int

	// defer 语句：在函数栈帧返回前插入
	defer close(ports)
	defer close(res)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, res)
	}

	go func() {
		for port := startPort; port <= endPort; port++ {
			ports <- port
		}
	}()

	for i := startPort; i <= endPort; i++ {
		p := <-res
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	sort.Ints(openPorts)
	for _, v := range openPorts {
		fmt.Printf("[+] %d open\n", v)
	}
}
