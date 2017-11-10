package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var connections []net.Conn

func main() {

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

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
	fmt.Println(len(connections))
}
