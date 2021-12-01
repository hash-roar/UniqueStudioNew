package main

import (
	"log"
	remoteserver "mysock5/remote_server"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	conf := remoteserver.RemoteServerConf{LisenAddr: tcpAddr}
	server, err := remoteserver.NewRemoteServer(&conf)
	if err != nil {
		log.Fatal(err)
	}
	server.ListenLocalServer()
	server.AcceptLocalConn()
}
