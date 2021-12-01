package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	const name = "user1"
	const pass = "pass1"
	buf := make([]byte, 512)
	message := make([]byte, 0)
	message = append(message, 0x01, 0x01, 0x01)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	conn.Write(message)
	length, err := io.ReadFull(conn, buf[:2])
	if length != 2 {
		log.Println(length)
		log.Println(err)
	}
	fmt.Println(buf[:2])
	message = make([]byte, 0)
	message = append(message, 0x01, byte(len(name)))
	message = append(message, []byte(name)...)
	message = append(message, byte(len(pass)))
	message = append(message, []byte(pass)...)
	conn.Write(message)
	length, err = io.ReadFull(conn, buf[0:2])
	if err != nil {
		log.Println(length)
		log.Println(err)
		return
	}
	if buf[1] == 0x00 {
		fmt.Println("验证成功")
	}
	

}
