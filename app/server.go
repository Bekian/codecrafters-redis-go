package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(">")

	// Read a line of text (until Enter is pressed)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	userInput = strings.TrimSpace(userInput)
	if userInput != "PING" {
		fmt.Println("Wrong input")
		os.Exit(1)
	} else {
		fmt.Println("PONG")
	}

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	fmt.Println(l)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
