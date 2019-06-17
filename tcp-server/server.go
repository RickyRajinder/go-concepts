package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

/**
Concurrent TCP server
Server will start a new goroutine for each new request
Tutorial from: https://opensource.com/article/18/5/building-concurrent-tcp-server-go
 */


func handleConnection(c net.Conn){
	read := bufio.NewReader(os.Stdin)
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
		fmt.Print("Message received from client: " + string(netData))
		fmt.Println("Send response back to client: ")
		msg, _ := read.ReadString('\n')
		c.Write([]byte(string(msg)))
		fmt.Println("Waiting on response from client...")
	}
	c.Close()
}

func serve(){
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a port number.")
		return
	}

	PORT := ":" + args[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port " + args[1])
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

func main() {
	serve()
}