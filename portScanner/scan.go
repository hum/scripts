package portScanner

import (
	"fmt"
	"net"
	"sort"
)

func scanner(addr string, ports, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", addr, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}

func Scan(portMax int, address string, goroutinesLimit int) []int {
	ports := make(chan int, goroutinesLimit)
	results := make(chan int)
	var open []int

	for i := 0; i < cap(ports); i++ {
		go scanner(address, ports, results)
	}

	go func() {
		for i := 0; i <= portMax; i++ {
			ports <- i
		}
	}()

	for i := 0; i < portMax; i++ {
		port := <-results
		if port != 0 {
			open = append(open, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(open)

	return open
}
