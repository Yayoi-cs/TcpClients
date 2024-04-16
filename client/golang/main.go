package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type client struct {
	addr string
}

func newClient(addr string) *client {
	return &client{addr: addr}
}

func main() {
	c := newClient("localhost:8080")
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		response := string(buf[:n])
		fmt.Println("Response from Server : localhost:8080", response)
		if strings.Contains(response, "FLAG") {
			break
		} else if !strings.Contains(response, "??") {
			continue
		}
		str := strings.Split(response, " ")
		var parts string
		for _, part := range str {
			if strings.Contains(part, "+") {
				parts = part
				break
			}
		}
		nums := strings.Split(parts, "+")
		result := 0
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			result += num
		}
		fmt.Println("My response is", result)
		fmt.Println("In byte : ", []byte(strconv.Itoa(result)))
		conn.Write([]byte(strconv.Itoa(result)))
	}
}
