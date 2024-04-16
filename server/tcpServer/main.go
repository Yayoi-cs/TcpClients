package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	defer l.Close()
	if err != nil {
		fmt.Println("Error while listening")
		return
	}
	fmt.Println("Listening on localhost:8080")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error while Accepting")
			return
		}
		go tcpHandle(conn)
	}
}

func tcpHandle(conn net.Conn) {
	defer conn.Close()
	for i := 0; i < 10; i++ {
		num1 := rand.Intn(10000000)
		num2 := rand.Intn(10000000)
		payload := "No " + strconv.Itoa(i)
		payload += " : " + strconv.Itoa(num1) + "+" + strconv.Itoa(num2) + " = ??"
		answer := num1 + num2
		conn.Write([]byte(payload))
		startTime := time.Now()
		buf := make([]byte, 16)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error while reading")
			return
		}
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		resultString := "Time :" + elapsedTime.String() + "\n"
		if endTime.Sub(startTime).Seconds() > 2 {
			conn.Write([]byte("Sorry It's too late.\n"))
			return
		}
		correct := binary.LittleEndian.Uint32(buf) == uint32(answer)
		correct = correct || binary.BigEndian.Uint32(buf) == uint32(answer)
		tmp := ""
		userAns := 0
		if !correct {
			for _, c := range buf {
				if c == 0 {
					break
				}
				tmp += string(rune(c))
			}
			userAns, err = strconv.Atoi(tmp)
		}
		correct = correct || userAns == answer
		if err == nil && correct {
			conn.Write([]byte("Correct! " + resultString))
		} else {
			conn.Write([]byte("Failed!"))
			fmt.Println(buf, err)
			fmt.Println("Answer : ", answer)
			fmt.Println("Little : ", binary.LittleEndian.Uint32(buf))
			fmt.Println("Big : ", binary.BigEndian.Uint32(buf))
			fmt.Println("UserAns : ", userAns)
			return
		}
	}
	conn.Write([]byte("Congratulations!!! Here is your flag: FLAG{TCP_TESTING_IN_MANY_LANGUAGE}"))
	return
}
