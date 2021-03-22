package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("host", "localhost", "Host")
	port = flag.Int("port", 3090, "Port")
)

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go Broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(con net.Conn) {
	defer con.Close()
	message := make(chan string)
	go MessageWrite(con, message)
	clientName := con.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to server %s\n", clientName)
	messages <- fmt.Sprint("New client connected: ", clientName)
	incomingClients <- message
	inputMessage := bufio.NewScanner(con)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
	}
	leavingClients <- message
	messages <- fmt.Sprintf("%s left the chat", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}
