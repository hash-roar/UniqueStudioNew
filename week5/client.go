package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err) 
	inputReader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 2048)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err1 := conn.Write([]byte(input))
		checkError(err1)
		length, err2 := conn.Read(buffer)
		checkError(err2)
		if length == 0 {
			fmt.Println("server close connection")
		}
		fmt.Println(string(buffer[0:length]))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
