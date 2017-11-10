package main

import (
	"log"
	"net"
)

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

		conn.Write([]byte("Hello World!\n"))

		conn.Close()
	}
}
