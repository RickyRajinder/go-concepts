package main

import "net"
import "fmt"
import "bufio"
import "os"

var PORT = ""

func askforConn(){
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please enter a port number.")
		return
	}
	PORT = ":" + args[1]
}

func dialServer(){
	conn, _ := net.Dial("tcp", "localhost"+PORT)
	for {
		read := bufio.NewReader(os.Stdin)
		fmt.Println("Send message to server: ")
		msg, _ := read.ReadString('\n')
		_, _ = fmt.Fprintf(conn, msg+"\n")
		fmt.Println("Waiting on response from server...")
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Response received from server: " + response)
	}
}

func main(){
	askforConn()
	dialServer()
}