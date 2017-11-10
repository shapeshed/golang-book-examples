package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Message struct {
	conn    net.Conn
	message []byte
}

var connections []net.Conn
var messages = make(chan Message)

func main() {

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	go messageChannel()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		connections = append(connections, conn)
		fmt.Println(len(connections))

		go handleRequest(conn)

	}

}

func messageChannel() {
	for {
		select {

		case message := <-messages:
			broadcastMessage(&message)
		}
	}
}

func handleRequest(conn net.Conn) {
	for {

		message := make([]byte, 1024)

		_, err := conn.Read(message)
		if err != nil {
			if err == io.EOF {
				removeConn(conn)
				conn.Close()
				return
			}
			log.Fatal(err)
		}

		m := Message{
			conn:    conn,
			message: message,
		}

		messages <- m
	}

}

func broadcastMessage(m *Message) {
	for _, conn := range connections {
		if m.conn != conn {
			_, err := conn.Write(m.message)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func removeConn(conn net.Conn) {
	var i int
	for i = range connections {
		if connections[i] == conn {
			break
		}
	}
	connections = append(connections[:i], connections[i+1:]...)
}
