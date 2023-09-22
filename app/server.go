package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// start log
	fmt.Println("Logs from your program will appear here!")
	// listening on port :6379 using tcp
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	//fmt.Println(l)
	// log port bind failure
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	// create connection from the listener
	connection, err := l.Accept()
	// log connection failure
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	// create a reader on the connection to accept input
	reader := bufio.NewReader(connection)
	// read input until newline character (or enter key)
	message, err := reader.ReadString('\n')
	// log read failure
	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
		return
	}
	// log input recieved and send response
	fmt.Println("Received:", message)
	connection.Write([]byte("+PONG\r\n"))

}

/* function to handle connecitons dynamically
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
		return
	}
	fmt.Println("Received:", message)
	conn.Write([]byte("+PONG\r\n"))

}
*/
