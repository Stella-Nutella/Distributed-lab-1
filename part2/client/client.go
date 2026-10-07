package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println("> ", msg)
	}

}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin) //takes user inputs from commandline

	for {
		fmt.Println("Enter text: ")
		msg, _ := reader.ReadString('\n')
		if msg == "/exit\n" { //exit command
			break
		} else {
			fmt.Fprintln(conn, msg) //message print to server
		}
	}

}

func main() {
	// Get the server address and port from the commandline arguments.
	addressPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	conn, _ := net.Dial("tcp", *addressPtr) //connection

	go read(conn)
	write(conn)
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
