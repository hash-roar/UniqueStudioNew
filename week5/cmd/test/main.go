package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	const name = "user1"
	const pass = "pass1"
	const domain = "150.158.91.131"
	const port = "80"
	buf := make([]byte, 512)
	message := make([]byte, 0)
	message = append(message, 0x05, 0x01, 0x02)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
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
	log.Println(message)
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
	message = make([]byte, 0)
	message = append(message, 0x05, 0x01, 0x00, 0x03, byte(len(domain)))
	message = append(message, []byte(domain)...)
	desPort, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		log.Println(err)
	}
	portBuf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(portBuf, binary.BigEndian, desPort)
	desBites := portBuf.Bytes()
	message = append(message, desBites[len(desBites)-2:]...)
	conn.Write(message)

	length, err = io.ReadFull(conn, buf[:4])
	if err != nil {
		log.Println(buf[:length])
		log.Println(err)
	}
	fmt.Println(buf[:length])
	if buf[1] == 0x00 {
		log.Println("socks5 信道建立成功")
	}
	io.ReadFull(conn, buf[:6])
	message = make([]byte, 0)
	getString := "GET / HTTP/1.1"
	message = append(message, []byte(getString)...)
	message = append(message, '\r', '\n', '\n')
	conn.Write(message)
	length, err = conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buf[:length]))
}
