package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

/**
Concurrent TCP server
Server will start a new goroutine for each new request
Tutorial from: https://opensource.com/article/18/5/building-concurrent-tcp-server-go
 */

func handleConnection(c net.Conn){
	fmt.Printf("Serving address: %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		res := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(res)))
	}
	c.Close()
}


func random() int {
	return rand.Intn(99) + 1
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a port number")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}