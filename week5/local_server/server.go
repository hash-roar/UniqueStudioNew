package localserver

import (
	"errors"
	"fmt"
	"log"
	sockutils "mysock5/sock_utils"
	"net"
)

type ClientConn struct {
	localConnectAddr  net.Addr
	localConn         *net.Conn
	remoteConnentAddr net.Addr
	remoteConn        net.Conn
}

type LocalServerConfig struct {
	LocalListenAddr  *net.TCPAddr
	RemoteServerAddr *net.TCPAddr
}

type LocalServer struct {
	Conf                LocalServerConfig
	RemoteServerConnect net.Conn
	LocalListener       *net.TCPListener
}

func NewServer(conf *LocalServerConfig) (*LocalServer, error) {
	if conf.RemoteServerAddr == nil {
		return nil, errors.New("config error")
	}
	return &LocalServer{Conf: *conf}, nil
}

func (s *LocalServer) Listenlocal() {

	listener, err := net.ListenTCP("tcp", s.Conf.LocalListenAddr)
	if err != nil {
		log.Fatal("listen local port error")
	}
	s.LocalListener = listener
}

func (s *LocalServer) ReceiveLocalConn() {
	defer s.LocalListener.Close()

	for {
		clientConn, err := s.LocalListener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// ? socket read
		go s.clientConnHandler(clientConn)

	}

}

func (s *LocalServer) clientConnHandler(conn *net.TCPConn) {
	defer conn.Close()

	remoteServerConn, err := s.connectRemoteServer()
	if err != nil {
		log.Println(err)
		return
	}
	defer remoteServerConn.Close()
	remoteConn, err := sockutils.NewConfusionSock(remoteServerConn)
	if err != nil {
		return
	}
	s.Forwardata(conn, remoteConn)
}

func (s *LocalServer) Forwardata(src *net.TCPConn, des *sockutils.ConfusedSocket) {
	exitChan := make(chan bool, 1)
	go func(src *net.TCPConn, des *sockutils.ConfusedSocket, exit chan bool) {
		buf := make([]byte, 4096)
		// defer src.Close()
		// log.Println("return ")
		// defer des.Close()
		for {
			length, err := src.Read(buf)
			if err != nil {
				log.Print(err)
				exit <- true
				return
			}
			fmt.Println(buf[:length])
			des.Write(buf[:length])
		}
	}(src, des, exitChan)
	go func(src *net.TCPConn, des *sockutils.ConfusedSocket, exit chan bool) {
		buf := make([]byte, 4096)
		// defer src.Close()
		// log.Println("return ")
		// defer des.Close()
		for {
			length, err := des.Read(buf)
			if err != nil {
				log.Print(err)
				exit <- true
				return
			}
			fmt.Println(buf[:length])
			des.Write(buf[:length])
		}
	}(src, des, exitChan)
	<-exitChan
}

func (s *LocalServer) connectRemoteServer() (*net.TCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, s.Conf.RemoteServerAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
