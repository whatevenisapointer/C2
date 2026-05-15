package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func sendCommands(conn net.Conn) {
	command := "pwd"
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
