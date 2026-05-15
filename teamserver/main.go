package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal("Error", err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Error accepting connection")
		}

		fmt.Println("Implant connected: ", conn.RemoteAddr())
	}

}
