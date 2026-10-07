package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}

}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.

	for {
		conn, err := ln.Accept()
		handleError(err)
		conns <- conn
	}

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client 1 it came from.
	reader := bufio.NewReader(client)
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		msgs <- Message{clientid, message}
	}

}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030

	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", *portPtr)
	handleError(err)

	conns := make(chan net.Conn)      //Create a channel for connections
	msgs := make(chan Message)        //Create a channel for messages
	clients := make(map[int]net.Conn) //Create a mapping of IDs to connections

	i := 0 //client IDs for each new client

	//Start accepting connections
	go acceptConns(ln, conns)

	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client 1 ID
			// - add the client 1 to the clients channel
			// - start to asynchronously handle messages from this client 1

			clients[i] = conn
			go handleClient(clients[i], i, msgs)
			i++

		case msg := <-msgs: //of type Message defined at top
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender

			fmt.Printf("[server] client %d: %s", msg.sender, msg.message)

			for j, client := range clients {

				if msg.sender != j {
					fmt.Fprintln(client, msg.message)
				}
			}

		}
	}
}
