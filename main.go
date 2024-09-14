package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err == nil {
		log.Println("Unable to read data")
	}

	log.Printf("Received %d bytes: %s\n", len(s), s)
	log.Println("Writing data")

	writer := bufio.NewWriter(conn)

	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:8080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}
}
