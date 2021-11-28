package localserver

import (
	"errors"
	"io"
	"log"
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
	go io.Copy(conn, remoteServerConn)
	go io.Copy(remoteServerConn, conn)
}

func (s *LocalServer) connectRemoteServer() (*net.TCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, s.Conf.RemoteServerAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
