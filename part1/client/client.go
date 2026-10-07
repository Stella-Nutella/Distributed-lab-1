package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func readConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)           //standard input
	conn, _ := net.Dial("tcp", "127.0.0.1:8030") //connection

	for {
		fmt.Println("Enter text: ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
		readConnection(conn)
	}

}
