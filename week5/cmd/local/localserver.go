package main

import (
	"log"
	localserver "mysock5/local_server"
	"net"
)

func main() {
	localIp, err := net.ResolveIPAddr("tcp", "127.0.0.1")
	if err != nil {
		log.Fatal("parse ip error")
	}
	conf := localserver.LocalServerConfig{LocalListenAddr: &net.TCPAddr{IP: localIp.IP, Port: 808},
		RemoteServerAddr: &net.TCPAddr{IP: localIp.IP, Port: 8081}}
	server, err := localserver.NewServer(&conf)
	if err != nil {
		log.Fatal(err)
	}
	server.Listenlocal()
	server.ReceiveLocalConn()
}
