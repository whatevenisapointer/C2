package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatal("[-] Error", err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(" [-] Error accepting connection")
		}

		fmt.Println("Implant connected: ", conn.RemoteAddr())
		go sendCommands(conn) //only using go routine here for future compatibility multiple cons
	}

}

func sendCommands(conn net.Conn) {
	command := "dir"
	_, err := conn.Write([]byte(command))
	if err != nil {
		log.Fatal("[-] Error sending command", err)
	}
	defer conn.Close()

	fmt.Println("[+] Command sent successfully")
}
